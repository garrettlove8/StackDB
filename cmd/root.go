package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sdb",
	Short: "A stackable database for cloud native applications",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, from StackDB!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
