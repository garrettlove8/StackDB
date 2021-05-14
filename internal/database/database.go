package database

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type Database struct {
	Uuid        string       `json:"uuid"`
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	CTime       string       `json:"cTime"`
	MTime       string       `json:"mTime"`
	Collections []Collection `json:"collections"`
}

type DatabaseMeta struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	CTime string `json:"cTime"`
	MTime string `json:"mTime"`
}

// Create creates a new database using the database struct that it is passed.
func (db *Database) Create() (*Database, error) {
	if err := createDatabaseDir(db.Name); err != nil {
		return nil, err
	}

	file, err := createDbFile(db.Name)
	if err != nil {
		return nil, err
	}

	if err = writeDbFile(file, *db); err != nil {
		return nil, err
	}

	if err = createCollectionsDir(db.Name); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Database) Insert(colName string, data *Data) (*Data, error) {
	return data, nil
}

func (db *Database) Read() (*DatabaseMeta, error) {
	meta := DatabaseMeta{
		Uuid: db.Uuid,
		Name: db.Name,
		Type: db.Type,
	}
	return &meta, nil
}

func (db *Database) Edit(newDatabase *Database) (*DatabaseMeta, error) {
	db.Name = newDatabase.Name

	meta := DatabaseMeta{
		Uuid:  db.Uuid,
		Name:  db.Name,
		Type:  db.Type,
		CTime: db.CTime,
		MTime: db.MTime,
	}
	return &meta, nil
}

func (db *Database) Delete() error {
	return nil
}

func (db *Database) Load() (*Database, error) {
	data, err := readDbFile(db.Name)
	if err != nil {
		return nil, err
	}

	db, err = decodeDbFile(data)
	if err != nil {
		return nil, err
	}

	for _, v := range db.Collections {
		data, _ := readColFile(db.Name, v.Name)
		col, _ := decodeColFile(data)
		db.Collections = append(db.Collections, *col)
	}

	return db, nil
}

func (db *Database) Search(data *Data) (*[]Data, error) {
	dataSlice := make([]Data, 0)

	dataSlice[0] = Data{}

	return &dataSlice, nil
}

func createDatabaseDir(dbName string) error {
	if err := os.MkdirAll("./stackdb/data/"+dbName, 0777); err != nil {
		return err
	}

	return nil
}

func createCollectionsDir(dbName string) error {
	if err := os.MkdirAll("./stackdb/data/"+dbName+"/collections", 0777); err != nil {
		return err
	}

	return nil
}

func createDbFile(dbName string) (*os.File, error) {
	file, err := os.Create("./stackdb/data/" + dbName + "/database.json")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func readDbFile(dbName string) ([]byte, error) {
	contentBytes, err := ioutil.ReadFile("./stackdb/data/" + dbName + "/database.json")
	if err != nil {
		return nil, err
	}

	return contentBytes, nil
}

func decodeDbFile(data []byte) (*Database, error) {
	var database Database

	err := json.Unmarshal(data, &database)
	if err != nil {
		return nil, err
	}

	return &database, nil
}

func writeDbFile(file *os.File, db Database) error {
	// Get the base db json file and convert its contents into a byte array
	pwd, _ := os.Getwd()
	content, _ := ioutil.ReadFile(pwd + "/configs/" + os.Getenv("VERSION") + "/baseDbFile.json")
	contentBytes := []byte(content)

	// Convert byte array into instance of Database struct for manipulation using Go
	database, err := decodeDbFile(contentBytes)

	// Now that we have a instance of a database struct, we can assign the values we want
	database.Uuid = uuid.New().String()
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
