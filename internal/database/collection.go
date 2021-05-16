package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Collections are the same as those in many NoSQL database and are akin to tables in relational databases.
// They privde an easy and logical way to separate data with a database
type Collection struct {
	// Uuid (Universal Unique Identifier) is the ID for a collection.
	// This field is internally managed and included in a collection's meta data.
	Uuid string `json:"uuid"`

	// Name is the name of a collection. As the developer, you'll use this field often.
	Name string `json:"name"`

	// CTime (Creation Time) is the time at which a collection was created.
	// This field is internally managed and included in a collection's meta data.
	CTime string `json:"cTime"`

	// MTime (Modified Time) is the time at which a collection was last changed.
	// This field is internally managed and included in a collection's meta data.
	MTime string `json:"mTime"`

	// Data is the data held within a collection.
	Data map[string][]Data `json:"data"`
}

// CollectionMeta is the meta data representation of a collection.
// All fields within CollectionMeta follow the same rules, guidelines, and usage
// as they do in the Collection type.
type CollectionMeta struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	CTime string `json:"cTime"`
	MTime string `json:"mTime"`
}

// Create facilitates the creation of a new collection in a database.
// It's job is to create the necessary directories for the new collection,
// after which it handles updating the its database to account for itself.
func (c *Collection) Create(dbName string) error {
	fmt.Println("collection:Create:c: ", c)

	// Open database file
	pwd, _ := os.Getwd()
	content, err := ioutil.ReadFile(pwd + "/sdb/data/" + dbName + "/database.json")
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

	file, err := os.OpenFile(pwd+"/sdb/data/"+dbName+"/database.json", os.O_WRONLY, os.ModeAppend)
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

// Read provides access to a collection's meta data.
func (c *Collection) Read() (*CollectionMeta, error) {
	meta := CollectionMeta{
		Uuid:  c.Uuid,
		Name:  c.Name,
		CTime: c.CTime,
		MTime: c.MTime,
	}
	return &meta, nil
}

// Edit provides a way to edit a collection's meta data.
func (c *Collection) Edit() (*CollectionMeta, error) {
	meta := CollectionMeta{
		Uuid:  c.Uuid,
		Name:  c.Name,
		CTime: c.CTime,
		MTime: c.MTime,
	}
	return &meta, nil
}

// Delete provides a way to delete a collection from a database.
func (c *Collection) Delete() error {
	return nil
}

func readColFile(dbName string, colName string) ([]byte, error) {
	contentBytes, err := ioutil.ReadFile("./sdb/data/" + dbName + "/collections/" + colName + ".json")
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
	file, err := os.Create("./sdb/data/" + dbName + "/collections" + "/" + colName + ".json")
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
