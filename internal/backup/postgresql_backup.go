package backup

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jonalphabert/db-guard/internal/logger"
	"github.com/jonalphabert/db-guard/internal/models"
	"github.com/jonalphabert/db-guard/internal/shared"
)

func BackupPostgreSQL(ctx context.Context, config models.Config) error {
	backupDirectory, err := resolveBackupDirectory(config)
	if err != nil {
		return err
	}

	err = os.MkdirAll(backupDirectory, 0755)
	if err != nil {
		return fmt.Errorf("failed to create backup directory: %w", err)
	}

	fileName := buildBackupFileName(config.DatabaseConfig.DbName, "postgres")
	outputPath := filepath.Join(backupDirectory, fileName)

	logger.Info("Preparing PostgreSQL backup file at %s", outputPath)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create backup file: %w", err)
	}

	defer outputFile.Close()

	commandArguments := buildPostgresArguments(config)
	command := exec.CommandContext(ctx, "pg_dump", commandArguments...)

	command.Stdout = outputFile
	command.Stderr = os.Stderr
	command.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", config.DatabaseConfig.Password))

	logger.Info("Running pg_dump for database %s", config.DatabaseConfig.DbName)

	err = command.Run()
	if err != nil {
		logger.Error("PostgreSQL backup failed: %v", err)

		return fmt.Errorf("pg_dump command failed: %w", err)
	}

	logger.Success("PostgreSQL backup file created at %s", outputPath)

	return nil
}

func buildPostgresArguments(config models.Config) []string {
	port := strconv.Itoa(config.DatabaseConfig.Port)

	arguments := []string{
		"-h", config.DatabaseConfig.Host,
		"-p", port,
		"-U", config.DatabaseConfig.User,
		"-d", config.DatabaseConfig.DbName,
		"-F", "p",
	}

	return arguments
}

func resolveBackupDirectory(config models.Config) (string, error) {
	if config.BackupConfig.Dir != "" {
		return config.BackupConfig.Dir, nil
	}

	baseDirectory, err := shared.BaseDir()
	if err != nil {
		return "", fmt.Errorf("failed to resolve base directory: %w", err)
	}

	backupDirectory := filepath.Join(baseDirectory, "backups")

	return backupDirectory, nil
}

func buildBackupFileName(databaseName string, engineName string) string {
	currentTime := time.Now().UTC()
	timestamp := currentTime.Format("20060102-150405")

	fileName := fmt.Sprintf("%s-%s-%s.sql", databaseName, engineName, timestamp)

	return fileName
}
