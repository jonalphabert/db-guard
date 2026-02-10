/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"github.com/jonalphabert/db-guard/internal/config"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage DB Guard configuration",
	Long: `Manage DB Guard configuration settings.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config.ConfigLocation()
		return nil
	},
}

func init() {
	// rootCmd.AddCommand(ConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
