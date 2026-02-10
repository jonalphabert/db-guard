/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jonalphabert/db-guard/internal/logviewer"
)

var (
	logLevel string
	logTail  int
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Show DB Guard logs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return logviewer.Show(logLevel, logTail)
	},
}

func init() {
	logsCmd.Flags().StringVarP(&logLevel, "level", "l", "INFO", "Log level (DEBUG, INFO, WARN, ERROR)")
	logsCmd.Flags().IntVarP(&logTail, "tail", "t", 100, "Show last N lines")

	rootCmd.AddCommand(logsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
