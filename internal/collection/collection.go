package collection

import (
	"StackDB/internal/utils"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Collections are the same as those in many NoSQL database and are akin to tables in relational databases.
// They privde an easy and logical way to separate data with a database
type Collection struct {
	// Uuid (Universal Unique Identifier) is the ID for a collection.
	// This field is internally managed and included in a collection's meta data.
	Uuid string `json:"uuid"`

	// Name is the name of a collection. As the developer, you'll use this field often.
	Name string `json:"name"`

	// CTime (Created Time) is the time at which a collection was created.
	// This field is internally managed and included in a collection's meta data.
	CTime string `json:"cTime"`

	// UTime (Updated Time) is the time at which a collection was last changed.
	// This field is internally managed and included in a collection's meta data.
	UTime string `json:"mTime"`

	// Data is the data held within a collection.
	Data map[string]Data `json:"data"`
}

// Create facilitates the creation of a new collection in a database.
// It's job is to create the necessary directories for the new collection,
// after which it handles updating the its database to account for itself.
func NewCollection() *Collection {
	return &Collection{
		Uuid:  utils.GetUuid(),
		CTime: time.Now().String(),
		UTime: time.Now().String(),
	}
}

// Create facilitates the creation of a new collection in a database.
// It's job is to create the necessary directories for the new collection,
// after which it handles updating the its database to account for itself.
func (c *Collection) Create() (*Collection, error) {
	fmt.Println("collection:Create:c: ", c)

	return nil, nil
}

// Read provides access to a collection's meta data.
func (c *Collection) Read() (*Collection, error) {
	meta := Collection{
		Uuid:  c.Uuid,
		Name:  c.Name,
		CTime: c.CTime,
		UTime: c.UTime,
	}
	return &meta, nil
}

// Edit provides a way to edit a collection's meta data.
func (c *Collection) Edit() (*Collection, error) {
	meta := Collection{
		Uuid:  c.Uuid,
		Name:  c.Name,
		CTime: c.CTime,
		UTime: c.UTime,
	}
	return &meta, nil
}

// Delete provides a way to delete a collection from a database.
func (c *Collection) Delete() error {
	return nil
}

// Delete provides a way to delete a collection from a database.
func (c *Collection) Persist() error {
	return nil
}

// Delete provides a way to delete a collection from a database.
func (c *Collection) Load() (*Collection, error) {
	return nil, nil
}

// Delete provides a way to delete a collection from a database.
func (c *Collection) Insert(data *Data) (*Collection, error) {
	return nil, nil
}

func saveDbFile(file *os.File, col *Collection) error {
	// Convert database struct back to json
	databaseJson, err := json.Marshal(col)
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
