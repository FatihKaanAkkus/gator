package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	return write(*c)
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	data, err := fs.ReadFile(os.DirFS(homeDir), configFileName)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func write(cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s", homeDir, configFileName), data, 0644)
	if err != nil {
		return err
	}
	return nil
}
