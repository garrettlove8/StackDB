package cmd

import (
	"StackDB/internal/database"
	"StackDB/internal/setup"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var systemDatabase *database.Database
var activeDatabase *database.Database

var rootCmd = &cobra.Command{
	Use:   "sdb",
	Short: "A stackable database for cloud native applications",
	Run: func(cmd *cobra.Command, args []string) {
		if isSetup := setup.CheckSetup(); !isSetup {
			fmt.Println("StackDB has not been setup yet")
		}
	},
}

func Execute(active, system *database.Database) {
	activeDatabase = active
	systemDatabase = system

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
