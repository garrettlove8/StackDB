package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Collection struct {
	Id   string                 `json:"id"`
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

func LoadCollection(dbName string) (*Database, error) {
	var database Database
	dataBytes := readDbFile(dbName)

	err := json.Unmarshal(dataBytes, &database)
	if err != nil {
		return nil, err
	}

	return &database, nil
}

func (c *Collection) CreateCollection(dbName string) error {
	fmt.Println("collection:Create:c: ", c)

	// Open database file
	pwd, _ := os.Getwd()
	content, err := ioutil.ReadFile(pwd + "/stackdb/data/" + dbName + "/database.json")
	if err != nil {
		fmt.Println("Unable to open database file: ", err)
		return err
	}
	contentBytes := []byte(content)

	// Unmarshal database file to Database struct
	var dbFileStruct Database
	if err := json.Unmarshal(contentBytes, &dbFileStruct); err != nil {
		return err
	}

	// Add collection to database collections slice
	dbFileStruct.Collections = append(dbFileStruct.Collections, c.Name)
	fmt.Printf("dbFileStruct: %s\n", dbFileStruct.Collections)

	file, err := os.OpenFile(pwd+"/stackdb/data/"+dbName+"/database.json", os.O_WRONLY, os.ModeAppend)
	err = saveDbFile(file, dbFileStruct)
	if err != nil {
		fmt.Printf("saveDbFile: %v\n", err)
		return err
	}

	// Create new collection file in proper database directory

	file, err = createColFile(dbName, c.Name)
	if err != nil {
		return err
	}

	// Convert Collection struct to JSON
	colJson, err := json.Marshal(c)
	if err != nil {
		return err
	}

	// Convert json to bytes
	colBytes := []byte(colJson)

	// Write bytes to collection file
	err = writeColFile(file, colBytes)

	// Save file to disk
	file.Sync()

	return nil
}

func readColFile(dbName string, colName string) []byte {
	fmt.Println("./stackdb/data/" + dbName + "/" + colName + ".json")
	contentBytes, _ := ioutil.ReadFile("./stackdb/data/" + dbName + "/collections/" + colName + ".json")

	return contentBytes
}

func encodeColFile(data []byte) (*Collection, error) {
	var collection Collection

	err := json.Unmarshal(data, &collection)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}

func saveDbFile(file *os.File, db Database) error {
	// Convert database struct back to json
	databaseJson, err := json.Marshal(db)
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

	file.Sync()

	return nil
}

func createColFile(dbName string, colName string) (*os.File, error) {
	file, err := os.Create("./stackdb/data/" + dbName + "/collections" + "/" + colName + ".json")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func writeColFile(file *os.File, data []byte) error {
	_, err := file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
