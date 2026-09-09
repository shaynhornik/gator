package config

import (
	"fmt"
	"os"
	"encoding/json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath () (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Failed to find home directory: %w",err)
	}
	configPath := home + "/" + configFileName
	return configPath, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Read () (Config, error)  {
	cfg := Config{}
	path, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	return write(*cfg)
}
