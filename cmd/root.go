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

func Execute(active, system *database.Database) error {
	if isSetup := setup.CheckSetup(); isSetup {
		wantedDb := database.Database{
			Name: "stackdb",
		}

		db, err := wantedDb.Load()
		if err != nil {
			return fmt.Errorf("Unable to load system database")
		}

		activeDatabase = active
		systemDatabase = db
		system = db
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return nil
}
