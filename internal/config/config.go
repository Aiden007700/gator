package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url            string
	Current_user_name string
}

func Read() (Config, error) {
	buff, err := os.ReadFile(configFileName)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(buff, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) SetUser(user string) error {
	c.Current_user_name = user

	if err := writeConfigToJSONFile(c); err != nil {
		return err
	}
	return nil
}

func writeConfigToJSONFile(c *Config) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	if err := os.WriteFile(configFileName, data, 0644); err != nil {
		return err
	}
	return nil
}
