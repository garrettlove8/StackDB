package cmd

import (
	"StackDB/internal/collections"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(collectionCmd)
	collectionCmd.AddCommand(createCollectionCmd)
}

var collectionCmd = &cobra.Command{
	Use:   "collections",
	Short: "The collection command allows you to interact with any of the collections in the active database.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the collection command")
		fmt.Println("Create collection : database : active : ", activeCollection)
		fmt.Println("Create collection : database : system : ", systemCollection)
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
			return fmt.Errorf("not enough arguments")
		}

		newCollection, _ := collections.NewCollection(args...)
		newCollection.Name = args[0]

		// _, err := newCollection.Create()
		// if err != nil {
		// 	return fmt.Errorf("Unable to create database: %v", err)
		// }

		// err = activeCollection.Persist()
		// if err != nil {
		// 	return fmt.Errorf("unable to persist database: %v\n Error: %v", activeCollection.Name, err)
		// }

		fmt.Printf("Database created:\n\tName: %v\n", args[0])

		return nil
	},
}
