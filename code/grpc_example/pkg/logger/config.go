package logger

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

const configPath = "config/zap_config.json"

func getZapConfig() (*zap.Config, error) {
	configFromFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	cfg := &zap.Config{}
	if err := json.Unmarshal(configFromFile, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
