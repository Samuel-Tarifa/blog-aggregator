package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config

	if err := json.Unmarshal(raw, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(home, configFileName)
	return path, nil
}

func (c Config) SetUser(name string) error {
	c.CurrentUserName = name
	err := write(c)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	b, err := json.MarshalIndent(cfg,"","  ")
	if err != nil {
		return err
	}
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	return os.WriteFile(path,b,0600)
}
