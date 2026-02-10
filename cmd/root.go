/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/jonalphabert/db-guard/cmd/config"
	"github.com/jonalphabert/db-guard/internal/logger"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "db-guard",
	Short: "DB Guard is a simple CLI tool for backing up databases and verifying that backups can actually be restored.",
	Long: `DB Guard is a developer-focused CLI tool designed to make database backups reliable, not just automated.
Most backup tools stop at creating dump files. dbkeeper goes one step further by testing the restore process, so you know your backups actually work before you need them.

The tool runs close to your database, uses official client tools like pg_dump and mysqldump, and keeps everything simple, transparent, and scriptable. No hidden magic, no cloud lock-in — just predictable backups with clear logs and results.

DB Guard is built for developers and small teams who want:
- simple configuration
- clear logs
- verified backups
- and full control over their infrastructure`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.db-guard.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cobra.OnInitialize(initLogger)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(config.ConfigCmd)
}

func initLogger() {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	logPath := filepath.Join(home, ".dbkeeper", "logs", "dbkeeper.log")
	_ = logger.Init(logPath)
}

