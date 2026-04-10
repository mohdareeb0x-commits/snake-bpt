/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"snake-bpt/logic"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// fastapiCmd represents the fastapi command
var fastapiCmd = &cobra.Command{
	Use:   "fastapi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		host, err := cmd.Flags().GetString("host")
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}
		
		if !logic.ValidateIP(host) {
			os.Exit(1)
		}

		port, err := cmd.Flags().GetString("port")
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		if !logic.ValidatePort(port) {
			os.Exit(1)
		}

		err = logic.AddFastAPI(host, port)
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		pterm.Success.Println("FastAPI project initialized")
	},
}

func init() {
	createCmd.AddCommand(fastapiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fastapiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fastapiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	fastapiCmd.Flags().String("host", "0.0.0.0", "IP address to bind")
	fastapiCmd.Flags().String("port", "8000", "Port to bind")
}
