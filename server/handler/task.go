package handler

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

type TaskStatus string

const (
	TaskPending   TaskStatus = "pending"
	TaskCompleted TaskStatus = "completed"
	TaskFailed    TaskStatus = "failed"
)

type Task struct {
	ID      string
	Command string
	Args    []string
	Status  TaskStatus
	Result  string
}

type TaskManager struct {
	db *sql.DB
}

func NewTaskManager(db *sql.DB) *TaskManager {
	return &TaskManager{
		db: db,
	}
}

func (tm *TaskManager) AddTask(agentID string, cmd string, args []string) *Task {
	taskID := "task_" + time.Now().Format("20060102150405") // ponytail: simple timestamp ID
	
	argsStr := strings.Join(args, ",")
	_, err := tm.db.Exec("INSERT INTO tasks (id, agent_id, command, args, status, created_at) VALUES (?, ?, ?, ?, ?, ?)",
		taskID, agentID, cmd, argsStr, TaskPending, time.Now())
	if err != nil {
		// ponytail: handle error
	}

	return &Task{
		ID:      taskID,
		Command: cmd,
		Args:    args,
		Status:  TaskPending,
	}
}

func (tm *TaskManager) GetPendingTask(agentID string) (*Task, error) {
	var t Task
	var argsStr string
	
	err := tm.db.QueryRow("SELECT id, command, args, status FROM tasks WHERE agent_id = ? AND status = ? LIMIT 1",
		agentID, TaskPending).Scan(&t.ID, &t.Command, &argsStr, &t.Status)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no tasks pending")
		}
		return nil, err
	}

	t.Args = strings.Split(argsStr, ",")
	return &t, nil
}

func (tm *TaskManager) CompleteTask(agentID, taskID, result string) {
	_, _ = tm.db.Exec("UPDATE tasks SET status = ?, result = ? WHERE id = ? AND agent_id = ?",
		TaskCompleted, result, taskID, agentID)
}

func (tm *TaskManager) GetTasksForAgent(agentID string) ([]*Task, error) {
	rows, err := tm.db.Query("SELECT id, command, args, status, result FROM tasks WHERE agent_id = ? ORDER BY created_at DESC", agentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		var t Task
		var argsStr string
		if err := rows.Scan(&t.ID, &t.Command, &argsStr, &t.Status, &t.Result); err != nil {
			continue
		}
		t.Args = strings.Split(argsStr, ",")
		tasks = append(tasks, &t)
	}
	return tasks, nil
}
