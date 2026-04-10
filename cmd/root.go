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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.snake-bpt.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


