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

// cobraCmd represents the cobra command
var cobraCmd = &cobra.Command{
	Use:   "cobra",
	Short: "Generate a Cobra CLI project scaffold",
	Long: `Cobra creates a new Cobra-based CLI project scaffold.
Use the --name flag to specify the CLI executable name.

Examples:
  snake-bpt create cobra --name my-cli
  snake-bpt create cobra --name mytool`,
	Run: func(cmd *cobra.Command, args []string) {

		cliName, err := cmd.Flags().GetString("name")
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		err = logic.AddCobra(cliName)
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		pterm.Success.Println("CLI initialized with name:", cliName)
		
	},
}

func init() {
	createCmd.AddCommand(cobraCmd)
	
	cobraCmd.Flags().String("name", "my-cli", "To initialize cli name")
}
