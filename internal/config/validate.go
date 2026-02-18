package config

import (
	"fmt"

	"github.com/jonalphabert/db-guard/internal/models"
)

func ValidateConfigRule(config models.Config) error {
	invalidatedRule := false
	if config.DatabaseConfig.Host == "" {
		fmt.Println("database host is required")
		invalidatedRule = true
	}
	if config.DatabaseConfig.Port == 0 {
		fmt.Println("database port is required")
		invalidatedRule = true
	}
	if config.DatabaseConfig.User == "" {
		fmt.Println("database user is required")
		invalidatedRule = true
	}
	if config.DatabaseConfig.Password == "" {
		fmt.Println("database password is required")
		invalidatedRule = true
	}
	if config.DatabaseConfig.DbName == "" {
		fmt.Println("database name is required")
		invalidatedRule = true
	}

	if invalidatedRule {
		return fmt.Errorf("Invalid configuration")
	}

	fmt.Println("Configuration is valid")
	return nil
}

func ValidateConfigFile(configPath string) error {
	config, err := readConfig(configPath)
	if err != nil {
		return err
	}
	return ValidateConfigRule(config)
}

func ValidateConfig() error {
	configPath, err := ConfigLocation()
	if err != nil {
		return err
	}
	return ValidateConfigFile(configPath)
}