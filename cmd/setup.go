package cmd

import (
	"StackDB/internal/setup"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A stackable database for cloud native applications",
	Run: func(cmd *cobra.Command, args []string) {
		err := setup.Setup()
		if err != nil {
			fmt.Println(err)
		}
	},
}
