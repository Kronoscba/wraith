package agent

type Config struct {
	C2Addr     string `json:"c2_addr"`
	SharedKey  string `json:"shared_key"`
	AgentID    string `json:"agent_id"`
	IntervalSec int    `json:"interval_sec"`
}

func DefaultConfig() *Config {
	return &Config{
		C2Addr:     "127.0.0.1:8080",
		SharedKey:  "this-is-a-32-byte-key-for-aes-256",
		AgentID:    "wraith-agent",
		IntervalSec: 10,
	}
}
