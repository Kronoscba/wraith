package main

import (
	"encoding/json"
	"log"
	"os"
	"github.com/Kronoscba/wraith/server"
	"github.com/Kronoscba/wraith/server/config"
	"github.com/Kronoscba/wraith/server/listener"
)

func main() {
	// 1. Load Config
	cfg := config.DefaultConfig()
	if data, err := os.ReadFile("config.json"); err == nil {
		json.Unmarshal(data, cfg)
	}

	// 2. Initialize Server
	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	// 3. Setup Listeners
	tcpL := listener.NewTCPListener(cfg.ServerPort)
	srv.AddListener(tcpL)

	// 4. Start Listeners in background
	go func() {
		if err := srv.Start(); err != nil {
			log.Printf("Listener error: %v", err)
		}
	}()

	// 5. Start API Server (Blocking)
	if err := srv.StartAPI(cfg.APIPort); err != nil {
		log.Fatalf("API server failed: %v", err)
	}
}
