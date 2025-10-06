package singleton

import (
	"encoding/json"
	"os"
	"sync"
)

type ConfigurationManager struct {
	settings map[string]string
}

var instance *ConfigurationManager
var once sync.Once

func GetInstance() *ConfigurationManager {
	once.Do(func() {
		instance = &ConfigurationManager{settings: make(map[string]string)}
	})
	return instance
}

func (c *ConfigurationManager) LoadFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &c.settings)
}

func (c *ConfigurationManager) SaveToFile(path string) error {
	data, err := json.MarshalIndent(c.settings, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (c *ConfigurationManager) Set(key, value string) {
	c.settings[key] = value
}

func (c *ConfigurationManager) Get(key string) string {
	return c.settings[key]
}
