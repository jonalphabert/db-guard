package setup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jonalphabert/db-guard/internal/logger"
)

func Init() error {
	home, err := os.UserHomeDir()
	if err != nil {
		logger.Error("Error getting user home directory")
		return err
	}

	baseDir := filepath.Join(home, ".dbkeeper")
	logDir := filepath.Join(baseDir, "logs")
	backupDir := filepath.Join(baseDir, "backups")
	configPath := filepath.Join(baseDir, "config.yaml")

	// Create directories
	for _, dir := range []string{baseDir, logDir, backupDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			logger.Error("Error creating directory")
			return err
		}
	}

	// Create config file if not exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.WriteFile(configPath, defaultConfig(), 0644); err != nil {
			logger.Error("Error creating config file")
			return err
		}
		logger.Info("Config file created: %s", configPath)
		fmt.Println("Created config:", configPath)
	} else {
		logger.Info("Config already exists: %s", configPath)
	}

	logger.Success("dbkeeper initialized")
	fmt.Println("dbkeeper initialized successfully")
	return nil
}

func defaultConfig() []byte {
	return []byte(`
database:
  type: postgres
  host: localhost
  port: 5432
  name: example_db
  user: example_user
  password: example_password

backup:
  path: ~/.dbkeeper/backups
  retention: 3
`)
}
