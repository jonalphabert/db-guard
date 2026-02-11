package config

import (
	"github.com/jonalphabert/db-guard/internal/logger"
	"github.com/jonalphabert/db-guard/internal/setup"
)

func ConfigLocation() (string, error) {
	configPath, err := setup.ConfigPath()
	if err != nil {
		logger.Error("Error getting config path: %v", err)
		return "", err
	}
	return configPath, nil
}