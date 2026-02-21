package backup

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/jonalphabert/db-guard/internal/logger"
	"github.com/jonalphabert/db-guard/internal/models"
)

func BackupMySQL(ctx context.Context, config models.Config) error {
	backupDirectory, err := resolveBackupDirectory(config)
	if err != nil {
		return err
	}

	err = os.MkdirAll(backupDirectory, 0755)
	if err != nil {
		return fmt.Errorf("failed to create backup directory: %w", err)
	}

	fileName := buildBackupFileName(config.DatabaseConfig.DbName, "mysql")
	outputPath := filepath.Join(backupDirectory, fileName)

	logger.Info("Preparing MySQL backup file at %s", outputPath)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create backup file: %w", err)
	}

	defer outputFile.Close()

	commandArguments := buildMySQLArguments(config)
	command := exec.CommandContext(ctx, "mysqldump", commandArguments...)

	command.Stdout = outputFile
	command.Stderr = os.Stderr
	command.Env = append(os.Environ(), fmt.Sprintf("MYSQL_PWD=%s", config.DatabaseConfig.Password))

	logger.Info("Running mysqldump for database %s", config.DatabaseConfig.DbName)

	err = command.Run()
	if err != nil {
		logger.Error("MySQL backup failed: %v", err)

		return fmt.Errorf("mysqldump command failed: %w", err)
	}

	logger.Success("MySQL backup file created at %s", outputPath)

	return nil
}

func buildMySQLArguments(config models.Config) []string {
	port := strconv.Itoa(config.DatabaseConfig.Port)

	arguments := []string{
		"-h", config.DatabaseConfig.Host,
		"-P", port,
		"-u", config.DatabaseConfig.User,
		config.DatabaseConfig.DbName,
	}

	return arguments
}
