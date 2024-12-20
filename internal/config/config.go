package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DB_URL          string `json:"db_url,omitempty"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}

const filename = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}

	return filepath.Join(home, filename), nil
}

func Read() *Config {
	configPath, err := getConfigFilePath()

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	content, err := os.ReadFile(configPath)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	rv := &Config{}
	if err := json.Unmarshal(content, rv); err != nil {

		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	return rv
}

func (this *Config) SetUser(username string) error {

	this.CurrentUserName = username
	return this.write()
}

func (this *Config) write() error {
	data, err := json.Marshal(this)

	if err != nil {
		return err
	}

	configPath, err := getConfigFilePath()

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}

	err = os.WriteFile(configPath, data, os.FileMode(os.O_WRONLY))

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}
	fmt.Println("Config wrote")
	return nil
}
