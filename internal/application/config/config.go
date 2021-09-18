package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Host                 string `json:"host"`
	Port                 string `json:"port"`
	ConnectionType       string `json:"connection_type"`
	MaxClientConnections int    `json:"max_client_connections"`
	SkuLogPath           string `json:"sku_log_path"`
}

func NewConfig(configPath string) (*Config, error) {

	config := &Config{}

	file, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	d := json.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
