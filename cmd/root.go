/*
Copyright © 2026 snake-bpt mohdareeb0x@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "snake-bpt",
	Short: "Snake-BPT CLI tool for bootstrapping projects",
	Long: `Snake-BPT is a command-line tool to scaffold application templates.
Use the create command with subcommands such as cobra or fastapi to
generate a Cobra CLI project or a FastAPI service, respectively.

Examples:
  snake-bpt create cobra --name my-cli
  snake-bpt create fastapi --host 127.0.0.1 --port 8000`,
  
	Run: func(cmd *cobra.Command, args []string) { 
		cmd.Help()
	},
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
	
}


