/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"github.com/jonalphabert/db-guard/internal/config"
	"github.com/spf13/cobra"
)

var (
	showAsJson bool
	showPassword bool
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show DB Guard configuration",
	Long: `Show DB Guard configuration settings.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config.Show(showAsJson, showPassword)
		return nil
	},
}

func init() {
	showCmd.Flags().BoolVar(&showAsJson, "show-json", false, "Show configuration as JSON")
	showCmd.Flags().BoolVar(&showPassword, "show-password", false, "Show password in plain text")

	ConfigCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
