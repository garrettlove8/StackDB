package collections

import (
	"StackDB/internal/utils"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// Collections are the same as those in many NoSQL database and are akin to tables in relational databases.
// They privde an easy and logical way to separate data with a database
type Collection struct {
	// Uuid (Universal Unique Identifier) is the ID for a Collection.
	// This field is internally managed and included in a Collection's meta data.
	Uuid string `json:"uuid"`

	// Name is the name of a Collection. As the developer, you'll use this field often.
	Name string `json:"name"`

	// CTime (Created Time) is the time at which a Collection was created.
	// This field is internally managed and included in a Collection's meta data.
	CTime string `json:"cTime"`

	// UTime (Updated Time) is the time at which a Collection was last changed.
	// This field is internally managed and included in a Collection's meta data.
	UTime string `json:"mTime"`

	// Location is the directory path to where the Collection is stored on disk.
	// By default, this is managed internally by StackDB, however, if you
	// are building on top of StackDB this may be helpful to override.
	Location string `json:"location"`

	// Data is the data held within a Collection.
	Data map[string]Data `json:"data"`
}

// NewCollection facilitates the creation of a new Collection in a database.
// It's job is to create the necessary directories for the new Collection,
// after which it handles updating the its database to account for itself.
//
// Accepts positional arguments: name, uuid, location string.
//
// Note: To save the returned Collection to disk use the Persist method.
func NewCollection(args ...string) (*Collection, error) {
	if len(args) == 0 {
		return nil, errors.New("no name provided for new Collection")
	}

	newCollection := Collection{
		Uuid:     utils.GetUuid(),
		CTime:    time.Now().String(),
		UTime:    time.Now().String(),
		Location: os.Getenv("DEFAULT_DATA_LOCATION"),
	}

	newCollection.Name = args[0]

	if len(args) >= 2 {
		newCollection.Uuid = args[1]
	}

	if len(args) >= 3 {
		newCollection.Location = args[2]
	}

	return &newCollection, nil
}

// Read provides access to a Collection's meta data.
func (c *Collection) Read() (*Collection, error) {
	meta := Collection{
		Uuid:  c.Uuid,
		Name:  c.Name,
		CTime: c.CTime,
		UTime: c.UTime,
	}
	return &meta, nil
}

// Edit provides a way to edit a Collection's meta data.
func (c *Collection) Edit() (*Collection, error) {
	meta := Collection{
		Uuid:  c.Uuid,
		Name:  c.Name,
		CTime: c.CTime,
		UTime: c.UTime,
	}
	return &meta, nil
}

// Delete provides a way to delete a Collection from a database.
func (c *Collection) Delete() error {
	return nil
}

// Delete provides a way to delete a Collection from a database.
func (c *Collection) Persist(file *os.File) error {
	var colBytes bytes.Buffer

	enc := gob.NewEncoder(&colBytes)
	err := enc.Encode(c)
	if err != nil {
		fmt.Println("c.Load: could not persist collection: ", err)
		return err
	}

	file.Write(colBytes.Bytes())

	return nil
}

// Delete provides a way to delete a Collection from a database.
func (c *Collection) Load() (*os.File, *Collection, error) {
	homepath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Unable to load collection: ", err)
		return nil, nil, err
	}

	// TODO: Change path to database configuration path
	file, err := os.OpenFile(homepath+"/sdb/data/"+c.Name+".json", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("c.Load: could not load collection: ", err)
		return nil, nil, err
	}

	col := Collection{}
	var fileBytes []byte

	_, err = file.Read(fileBytes)
	if err != nil {
		fmt.Println("c.Load: could not load collection: ", err)
		return nil, nil, err
	}

	json.Unmarshal(fileBytes, &col)

	return file, &col, nil
}

// Delete provides a way to delete a Collection from a database.
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
