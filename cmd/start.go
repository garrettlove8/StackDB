package cmd

import (
	"StackDB/internal/shell"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "The start command starts the StackDB server",
	Run: func(cmd *cobra.Command, args []string) {
		shell.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
