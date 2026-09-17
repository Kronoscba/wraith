package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"github.com/Kronoscba/wraith/server/handler"
)

type API struct {
	Agents *handler.AgentHandler
	Tasks  *handler.TaskManager
}

func NewAPI(agents *handler.AgentHandler, tasks *handler.TaskManager) *API {
	return &API{
		Agents: agents,
		Tasks:  tasks,
	}
}

func (a *API) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/agents", a.handleAgents)
	mux.HandleFunc("/agents/", a.handleAgentDetail)
	mux.HandleFunc("/tasks", a.handleTasks)
	mux.HandleFunc("/tasks/agent/", a.handleAgentTasks)
}

func (a *API) handleAgents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	agents, err := a.Agents.ListAgents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(agents)
}

func (a *API) handleAgentDetail(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/agents/")
	agent, ok := a.Agents.GetAgent(id)
	if !ok {
		http.Error(w, "Agent not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(agent)
}

func (a *API) handleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		AgentID string   `json:"agent_id"`
		Command string   `json:"command"`
		Args    []string `json:"args"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := a.Tasks.AddTask(req.AgentID, req.Command, req.Args)
	json.NewEncoder(w).Encode(task)
}

func (a *API) handleAgentTasks(w http.ResponseWriter, r *http.Request) {
	agentID := strings.TrimPrefix(r.URL.Path, "/tasks/agent/")
	tasks, err := a.Tasks.GetTasksForAgent(agentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}
