package tests

import (
	"testing"
	"time"
	"github.com/Kronoscba/wraith/agent"
	"github.com/Kronoscba/wraith/agent/comms"
	"github.com/Kronoscba/wraith/server"
	"github.com/Kronoscba/wraith/server/config"
	"github.com/Kronoscba/wraith/server/listener"
)

func TestEndToEndFlow(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.DBPath = "test_wraith.db"
	srv, err := server.NewServer(cfg)
	if err != nil {
		t.Fatalf("Server init failed: %v", err)
	}

	tcpL := listener.NewTCPListener("127.0.0.1:9002")
	srv.AddListener(tcpL)

	go srv.Start()
	time.Sleep(500 * time.Millisecond)

	channel := comms.NewTCPChannel("127.0.0.1:9002")
	session := agent.NewSession("test-agent", channel, 1*time.Second, []byte(cfg.SharedKey))
	go session.Run()

	// Poll hasta 10s
	var found bool
	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		a, ok := srv.Agents.GetAgent("test-agent")
		if ok && (a.Active || a.ID == "test-agent") {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("Agent did not register")
	}
}
