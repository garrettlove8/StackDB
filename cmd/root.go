package cmd

import (
	"StackDB/internal/database"
	"StackDB/internal/setup"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var systemDatabase *database.Database
var activeDatabase *database.Database

var rootCmd = &cobra.Command{
	Use:   "sdb",
	Short: "A stackable database for cloud native applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		if isSetup := setup.CheckSetup(); !isSetup {
			return errors.New("StackDB has not been setup yet")
		}

		// TODO: Add onto this message
		fmt.Println("Welcome to StackDB!")

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Execute:systemDatabase: ", systemDatabase)
}
