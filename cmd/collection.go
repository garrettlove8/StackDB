package cmd

import (
	"StackDB/internal/database"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(collectionCmd)
	collectionCmd.AddCommand(createCollectionCmd)
}

var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "The collection command allows you to interact with any of the collections in the active database.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the collection command")
		fmt.Println("Create collection : database : active : ", activeDatabase)
		fmt.Println("Create collection : database : system : ", systemDatabase)
	},
}

var createCollectionCmd = &cobra.Command{
	Use:   "create",
	Short: "The collection create command allows you to create a new collection in the active database",
	Long: `The collection create command allows you to create a new collection in the active database.
	Databases are created using the values of positional arguments, as follows
	1. name
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("ARGS: ", args)

		if len(args) < 1 {
			return fmt.Errorf("Not enough arguments")
		}

		newCollection := database.NewCollection()
		newCollection.Name = args[0]

		err := newCollection.Create(activeDatabase)
		if err != nil {
			return fmt.Errorf("Unable to create database: %v", err)
		}

		err = activeDatabase.Persist()
		if err != nil {
			return fmt.Errorf("unable to persist database: %v\n Error: %v", activeDatabase.Name, err)
		}

		fmt.Printf("Database created:\n\tName: %v\n", args[0])

		return nil
	},
}