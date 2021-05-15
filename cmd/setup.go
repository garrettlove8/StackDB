package cmd

import (
	"StackDB/internal/database"
	"StackDB/internal/setup"
	"fmt"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A stackable database for cloud native applications",
	Run: func(cmd *cobra.Command, args []string) {
		err := setup.Setup()
		if err != nil {
			fmt.Println(err)
		}

		system := database.Database{
			Name: "system",
		}
		systemDatabase, _ = system.Load()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
