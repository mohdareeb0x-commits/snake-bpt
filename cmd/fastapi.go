/*
Copyright © 2026 snake-bpt mohdareeb0x@gmail.com
*/
package cmd

import (
	"os"
	"github.com/mohdareeb0x-commits/snake-bpt/logic"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// fastapiCmd represents the fastapi command
var fastapiCmd = &cobra.Command{
	Use:   "fastapi",
	Short: "Generate a FastAPI project scaffold",
	Long: `FastAPI creates a new Python FastAPI project scaffold.
Use the host and port flags to configure the application bind address.

Examples:
  snake-bpt create fastapi --host 0.0.0.0 --port 8000
  snake-bpt create fastapi --host 127.0.0.1 --port 8080`,
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

	fastapiCmd.Flags().String("host", "0.0.0.0", "IP address to bind")
	fastapiCmd.Flags().String("port", "8000", "Port to bind")
}
