/*
Copyright © 2026 snake-bpt mohdareeb0x@gmail.com
*/
package cmd

import (
	"fmt"

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
		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
