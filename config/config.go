package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

// Config represents the configuration settings.
type Config struct {
	Namespace string `yaml:"namespace"`
}

// SetConfigOption sets a configuration option.
func SetConfigOption(key, value string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	switch key {
	case "namespace":
		config.Namespace = value
	default:
		return fmt.Errorf("unknown configuration option: %s", key)
	}

	if err := SaveConfig(config); err != nil {
		return err
	}

	return nil
}

// GetConfigOptions returns all configuration options as a map.
func GetConfigOptions() map[string]string {
	config, err := LoadConfig()
	if err != nil {
		return nil
	}

	configOptions := map[string]string{
		"namespace": config.Namespace,
	}

	return configOptions
}

// LoadConfig loads the configuration from the config file.
func LoadConfig() (*Config, error) {
	configFile, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig saves the configuration to the config file.
func SaveConfig(config *Config) error {
	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		return err
	}

	return nil
}

// getConfigFilePath returns the path to the config file.
func getConfigFilePath() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".config", "jtnctl")
	configFile := filepath.Join(configDir, "config.yaml")

	return configFile, nil
}
