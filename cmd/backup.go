/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"context"

	"github.com/jonalphabert/db-guard/internal/backup"
	"github.com/jonalphabert/db-guard/internal/logger"
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Create a database backup using official client tools",
	Long: `Create a database backup based on the current configuration.

This command executes the appropriate database client tool 
(pg_dump for PostgreSQL or mysqldump for MySQL) 
and stores the generated dump file in the configured backup directory.

Retention policy will be applied automatically if configured.
Logs will be written to the dbkeeper log file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Starting database backup...")

		contextInstance := context.Background()

		err := backup.Execute(contextInstance)
		if err != nil {
			logger.Error("Backup command failed: %v", err)
			return err
		}

		fmt.Println("Database backup completed successfully")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
