package cmd

import (
	"github.com/jonalphabert/db-guard/internal/setup"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Interactive configuration setup",
	RunE: func(cmd *cobra.Command, args []string) error {
		return setup.Run()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
