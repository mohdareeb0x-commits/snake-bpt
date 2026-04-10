/*
Copyright © 2026 snake-bpt mohdareeb0x@gmail.com
*/
package cmd

import (
	"os"
	"snake-bpt/logic"

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cobraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cobraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cobraCmd.Flags().String("name", "my-cli", "To initialize cli name")
}
