package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Collection struct {
	Uuid  string            `json:"uuid"`
	Name  string            `json:"name"`
	CTime string            `json:"cTime"`
	MTime string            `json:"mTime"`
	Data  map[string][]Data `json:"data"`
}

type CollectionMeta struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	CTime string `json:"cTime"`
	MTime string `json:"mTime"`
}

func (c *Collection) Create(dbName string) error {
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
	db, err := decodeDbFile(contentBytes)
	if err != nil {
		return err
	}

	// Add collection to database collections slice
	db.Collections = append(db.Collections, *c)

	file, err := os.OpenFile(pwd+"/stackdb/data/"+dbName+"/database.json", os.O_WRONLY, os.ModeAppend)
	err = saveDbFile(file, *db)
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

func (c *Collection) Read() (*CollectionMeta, error) {
	meta := CollectionMeta{
		Uuid:  c.Uuid,
		Name:  c.Name,
		CTime: c.CTime,
		MTime: c.MTime,
	}
	return &meta, nil
}

func (c *Collection) Edit() (*CollectionMeta, error) {
	meta := CollectionMeta{
		Uuid:  c.Uuid,
		Name:  c.Name,
		CTime: c.CTime,
		MTime: c.MTime,
	}
	return &meta, nil
}

func (c *Collection) Delete() error {
	return nil
}

func readColFile(dbName string, colName string) ([]byte, error) {
	contentBytes, err := ioutil.ReadFile("./stackdb/data/" + dbName + "/collections/" + colName + ".json")
	if err != nil {
		return nil, err
	}

	return contentBytes, nil
}

func decodeColFile(data []byte) (*Collection, error) {
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
