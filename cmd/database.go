package cmd

import (
	"StackDB/internal/database"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.AddCommand(createDatabaseCmd)
	databaseCmd.AddCommand(useDatabaseCmd)
}

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "The database command allows you to interact with any of the databases currently in the system.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the database command")
	},
}

var createDatabaseCmd = &cobra.Command{
	Use:   "create",
	Short: "The database create command allows you to create a new database",
	Long: `The database create command allows you to create a new database.
	Databases are created using the values of positional arguments, as follows
	1. name
	2. type
		Options: keyValue
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("ARGS: ", args)

		if len(args) < 2 {
			return fmt.Errorf("Not enough arguments")
		}

		if args[0] == "stackdb" {
			return fmt.Errorf("The word \"stackdb\" is reserved for StackDB usage")
		}

		if args[1] != "keyValue" {
			return fmt.Errorf("Unsupported database type")
		}

		newDatabase := database.Database{
			Name:  args[0],
			Type:  args[1],
			CTime: time.Now().String(),
			MTime: time.Now().String(),
		}

		_, err := newDatabase.Create()
		if err != nil {
			return fmt.Errorf("Unable to create database: %v", err)
		}

		fmt.Printf("Database create.\n\tName: %v\n\tType: %v\n", args[0], args[1])

		// TODO:  At this point, the database is technically created. However, we
		// still need to add it to the system database's database collection so
		// we can keep track of it

		return nil
	},
}

var useDatabaseCmd = &cobra.Command{
	Use:   "use",
	Short: "The use command allows you to load a given database so you can use it.",
	Long:  "The use command allows you to load a given database so you can use it. Only one database can be in use at a time.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("ARGS: ", args)

		if len(args) < 1 {
			return fmt.Errorf("Not enough arguments")
		}

		wantedDb := database.Database{
			Name: args[0],
		}

		db, err := wantedDb.Load()
		if err != nil {
			return fmt.Errorf("Unable to use database %v", args[0])
		}

		activeDatabase = db

		fmt.Printf("Database ready for use: %v", activeDatabase)

		return nil
	},
}
