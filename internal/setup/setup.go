package setup

import (
	"fmt"
	"os"

	"github.com/jonalphabert/db-guard/internal/logger"
)

func Run() error {
	baseDir, err := BaseDir()
	if err != nil {
		logger.Error("Failed to get base dir: %v", err)
		return err
	}

	// ensure base dir exists
	if err = os.MkdirAll(baseDir, 0755); err != nil {
		logger.Error("Failed to create base dir: %v", err)
		return err
	}

	configPath, err := ConfigPath()
	if err != nil {
		logger.Error("Failed to get config path: %v", err)
		return err
	}

	// detect existing config
	if _, err = os.Stat(configPath); err == nil {
		fmt.Println("✔ Found existing config at", configPath)
		fmt.Println("⚠ Overwriting existing configuration")
		logger.Info("Overwriting existing config at %s", configPath)
	}

	// run interactive wizard
	input, err := RunWizard()
	if err != nil {
		logger.Error("Interactive setup failed: %v", err)
		return err
	}

	cfg := RenderConfig(input)

	if err := os.WriteFile(configPath, cfg, 0644); err != nil {
		logger.Error("Failed to write config file: %v", err)
		return err
	}

	logger.Info("Database type set to %s", input.DBType)
	logger.Info("Host set to %s", input.Host)
	logger.Info("Port set to %d", input.Port)
	logger.Info("Database name set to %s", input.DBName)
	logger.Info("User set to %s", input.User)
	logger.Info("Password has been set successfully")
	
	logger.Info("Retention set to %d days", input.Retention)
	logger.Info("Backup directory set to %s", input.BackupDir)
	logger.Success("Configuration saved to %s successfully", configPath)

	fmt.Println("✔ Configuration saved to", configPath)
	return nil
}
