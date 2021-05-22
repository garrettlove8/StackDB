package cmd

import (
	"StackDB/internal/database"
	"StackDB/internal/setup"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.AddCommand(useDatabaseCmd)
	databaseCmd.AddCommand(createDatabaseCmd)
}

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "The database command allows you to interact with any of the databases currently in the system.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the database command")
		fmt.Println("Create collection : database : active : ", activeDatabase)
		fmt.Println("Create collection : database : system : ", systemDatabase)
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
		// TODO: Figure out debugging strategy so this can be implemented properly
		fmt.Println("ARGS: ", args)
		fmt.Println("database:create:systemDB: ", systemDatabase)

		if len(args) < 2 {
			return fmt.Errorf("Not enough arguments")
		}

		if args[0] == "stackdb" {
			return fmt.Errorf("The word \"stackdb\" is reserved for StackDB usage")
		}

		if args[1] != "keyValue" {
			return fmt.Errorf("Unsupported database type")
		}

		if isSetup := setup.CheckSetup(); !isSetup {
			return fmt.Errorf("please run setup process before creating a new database")
		}

		newDatabase := database.NewDatabase()

		newDatabase.Name = args[0]
		newDatabase.Type = args[1]
		newDatabase.CTime = time.Now().String()
		newDatabase.MTime = time.Now().String()

		_, err := newDatabase.Create()
		if err != nil {
			return fmt.Errorf("Unable to create database: %v", err)
		}

		body := make(map[string][]byte)
		body["name"] = []byte(newDatabase.Name)

		newData := database.NewData()
		newData.CTime = time.Now().String()
		newData.MTime = time.Now().String()
		newData.Body = body

		_, err = systemDatabase.Insert("databases", newData)
		if err != nil {
			// TODO: Idealy if there is an error here the process should be undone automatically.

			return fmt.Errorf(`
			database has been created,
			however there was an error adding the new database to the tracking system: %v.
			It is recommended to delete the new database and fix the tracking issue before recreating it.`,
				err)
		}

		err = systemDatabase.Persist()
		if err != nil {
			// TODO: Idealy if there is an error here the process should be undone automatically.

			return fmt.Errorf(`
			database has been created,
			however there was an error persisting the new database to the tracking system: %v.
			It is recommended to delete the new database and fix the tracking issue before recreating it.`,
				err)
		}

		fmt.Printf("Database created:\n\tName: %v\n\tType: %v\n", args[0], args[1])

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

		activeDatabase, _ = wantedDb.Load()
		// if err != nil {
		// 	return fmt.Errorf("Unable to use database %v", args[0])
		// }

		fmt.Printf("Database ready for use: %v", activeDatabase)

		return nil
	},
}
