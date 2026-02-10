package config

import (
	"fmt"

	"github.com/jonalphabert/db-guard/internal/logger"
	"github.com/jonalphabert/db-guard/internal/setup"
)

func ConfigLocation() {
	configPath, err := setup.ConfigPath()
	if err != nil {
		logger.Error("Error getting config path: %v", err)
		return
	}
	fmt.Printf("Config Path: %s\n", configPath)
}