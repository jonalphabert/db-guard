package backup

import (
	"context"
	"fmt"
	"strings"

	"github.com/jonalphabert/db-guard/internal/logger"
	"github.com/jonalphabert/db-guard/internal/shared"
)

func Execute(ctx context.Context) error {
	configPath, err := shared.ConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get config path: %w", err)
	}

	config, err := shared.ReadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	databaseType := strings.ToLower(config.DatabaseConfig.Type)

	if databaseType == "postgres" || databaseType == "postgresql" {
		logger.Info("Starting PostgreSQL backup")

		err := BackupPostgreSQL(ctx, config)
		if err != nil {
			return fmt.Errorf("failed to backup PostgreSQL: %w", err)
		}

		logger.Success("PostgreSQL backup completed successfully")

		return nil
	}

	if databaseType == "mysql" {
		logger.Info("Starting MySQL backup")

		err := BackupMySQL(ctx, config)
		if err != nil {
			return fmt.Errorf("failed to backup MySQL: %w", err)
		}

		logger.Success("MySQL backup completed successfully")

		return nil
	}

	return fmt.Errorf("unsupported database type: %s", config.DatabaseConfig.Type)
}
