package database

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type Database struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Collections []string `json:"collections"`
}

// Create creates a new database using the database struct that it is passed.
func (db *Database) CreateDatabase() (*Database, error) {
	err := createDatabaseDir(db.Name)
	if err != nil {
		return nil, err
	}

	file, err := createDbFile(db.Name)
	if err != nil {
		return nil, err
	}

	err = writeDbFile(file, *db)
	if err != nil {
		return nil, err
	}

	err = createCollectionsDir(db.Name)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createDatabaseDir(dbName string) error {
	err := os.MkdirAll("./stackdb/data/"+dbName, 0777)
	if err != nil {
		return err
	}

	return nil
}

func createCollectionsDir(dbName string) error {
	err := os.MkdirAll("./stackdb/data/"+dbName+"/collections", 0777)
	if err != nil {
		return err
	}

	return nil
}

func createDbFile(name string) (*os.File, error) {
	file, err := os.Create("./stackdb/data/" + name + "/database.json")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func readDbFile(dbName string) []byte {
	contentBytes, _ := ioutil.ReadFile("./stackdb/data/" + dbName + "/database.json")

	return contentBytes
}

func writeDbFile(file *os.File, db Database) error {
	// Get the base db json file and convert its contents into a byte array
	pwd, _ := os.Getwd()
	content, _ := ioutil.ReadFile(pwd + "/configs/" + os.Getenv("VERSION") + "/baseDbFile.json")
	contentBytes := []byte(content)

	// Convert byte array into instance of Database struct for manipulation using Go
	var database Database
	if err := json.Unmarshal(contentBytes, &database); err != nil {
		return err
	}

	// Now that we have a instance of a database struct, we can assign the values we want
	database.Id = uuid.New().String()
	database.Name = db.Name
	database.Type = db.Type

	// Convert database struct back to json
	databaseJson, err := json.Marshal(database)
	if err != nil {
		return err
	}

	// Convert database json back to a byte array so it can be written to the database file
	databaseBytes := []byte(databaseJson)

	// Write database byte array to database file
	_, err = file.Write(databaseBytes)
	if err != nil {
		return err
	}

	return nil
}
