package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerPort  string `json:"server_port"`
	APIPort     string `json:"api_port"`
	SharedKey   string `json:"shared_key"`
	DBPath      string `json:"db_path"`
	CertFile    string `json:"cert_file"`
	KeyFile     string `json:"key_file"`
}

func DefaultConfig() *Config {
	return &Config{
		ServerPort: "8080",
		APIPort:    "8000",
		SharedKey:  "this-is-a-32-byte-key-for-aes-ok",
		DBPath:     "wraith.db",
		CertFile:   "server.crt",
		KeyFile:    "server.key",
	}
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) Save(path string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
