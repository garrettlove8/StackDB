package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "The start command starts the StackDB server",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			reader := bufio.NewReader((os.Stdin))
			text, _ := reader.ReadString('\n')
			fmt.Print(text)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func processInput(input string) {
	fmt.Print(input)
}
