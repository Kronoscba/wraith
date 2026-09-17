package handler

import (
	"database/sql"
	"fmt"
	"net"
	"sync"
	"time"
)

type Agent struct {
	ID       string
	Conn    net.Conn
	LastSeen time.Time
	OS       string
	Arch     string
	Active   bool
}

type AgentHandler struct {
	db *sql.DB
	activeAgents map[string]*Agent
	mu           sync.RWMutex
}

func NewAgentHandler(db *sql.DB) *AgentHandler {
	return &AgentHandler{
		db:           db,
		activeAgents: make(map[string]*Agent),
	}
}

func (h *AgentHandler) RegisterAgent(id string, conn net.Conn) *Agent {
	h.mu.Lock()
	defer h.mu.Unlock()

	// Upsert into DB
	_, err := h.db.Exec("INSERT INTO agents (id, last_seen) VALUES (?, ?) ON CONFLICT(id) DO UPDATE SET last_seen=excluded.last_seen", id, time.Now())
	if err != nil {
		fmt.Printf("DB Error RegisterAgent: %v\n", err)
	}

	agent := &Agent{
		ID:       id,
		Conn:    conn,
		LastSeen: time.Now(),
	}
	h.activeAgents[id] = agent
	return agent
}

func (h *AgentHandler) DeactivateAgent(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if agent, ok := h.activeAgents[id]; ok {
		agent.Active = false
		agent.Conn = nil
	}
}

func (h *AgentHandler) GetAgent(id string) (*Agent, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	agent, ok := h.activeAgents[id]
	return agent, ok
}

func (h *AgentHandler) ListAgents() ([]*Agent, error) {
	rows, err := h.db.Query("SELECT id, os, arch, last_seen FROM agents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agents []*Agent
	for rows.Next() {
		var a Agent
		if err := rows.Scan(&a.ID, &a.OS, &a.Arch, &a.LastSeen); err != nil {
			continue
		}
		agents = append(agents, &a)
	}
	return agents, nil
}
