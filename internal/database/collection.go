package database

import (
	"StackDB/internal/utils"
	"encoding/json"
	"fmt"
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
	Data map[string]Data `json:"data"`
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
func NewCollection() *Collection {
	newCol := Collection{
		Uuid: utils.GetUuid(),
	}

	return &newCol
}

// Create facilitates the creation of a new collection in a database.
// It's job is to create the necessary directories for the new collection,
// after which it handles updating the its database to account for itself.
func (c *Collection) Create(db *Database) error {
	fmt.Println("collection:Create:c: ", c)
	fmt.Println("collection:Create:db:Collections: ", db.Collections)

	// Add collection to database collections slice
	db.Collections = append(db.Collections, *c)

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
