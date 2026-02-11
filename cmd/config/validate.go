/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"github.com/jonalphabert/db-guard/internal/config"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		config.ValidateConfig()
		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(validateCmd)
}
