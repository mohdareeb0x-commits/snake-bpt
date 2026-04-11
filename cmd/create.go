/*
Copyright © 2026 snake-bpt mohdareeb0x@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Scaffold a new project template",
	Long: `Create is the parent command for project scaffolding.
Use a subcommand to generate a specific type of template.

Examples:
  snake-bpt create cobra --name my-cli
  snake-bpt create fastapi --host 0.0.0.0 --port 8000`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
