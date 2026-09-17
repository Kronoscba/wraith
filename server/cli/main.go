package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wraith-cli <command> [args...]")
		fmt.Println("Commands: agents, task <agent_id> <cmd>")
		return
	}

	cmd := os.Args[1]
	baseURL := "http://127.0.0.1:8000"

	switch cmd {
	case "agents":
		resp, err := http.Get(baseURL + "/agents")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))

	case "task":
		if len(os.Args) < 4 {
			fmt.Println("Usage: wraith-cli task <agent_id> <cmd>")
			return
		}
		agentID := os.Args[2]
		command := os.Args[3]
		
		payload := map[string]interface{}{
			"agent_id": agentID,
			"command":  command,
			"args":     strings.Fields(command)[1:],
		}
		jsonPayload, _ := json.Marshal(payload)
		
		resp, err := http.Post(baseURL+"/tasks", "application/json", bytes.NewBuffer(jsonPayload))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Task sent:", string(body))

	default:
		fmt.Println("Unknown command")
	}
}
