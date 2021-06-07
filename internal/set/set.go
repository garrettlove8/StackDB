package set

import (
	"StackDB/internal/utils"
	"encoding/json"
	"errors"
	"os"
	"time"
)

// Sets are the same as those in many NoSQL database and are akin to tables in relational databases.
// They privde an easy and logical way to separate data with a database
type Set struct {
	// Uuid (Universal Unique Identifier) is the ID for a Set.
	// This field is internally managed and included in a Set's meta data.
	Uuid string `json:"uuid"`

	// Name is the name of a Set. As the developer, you'll use this field often.
	Name string `json:"name"`

	// CTime (Created Time) is the time at which a Set was created.
	// This field is internally managed and included in a Set's meta data.
	CTime string `json:"cTime"`

	// UTime (Updated Time) is the time at which a Set was last changed.
	// This field is internally managed and included in a Set's meta data.
	UTime string `json:"mTime"`

	// Location is the directory path to where the set is stored on disk.
	// By default, this is managed internally by StackDB, however, if you
	// are building on top of StackDB this may be helpful to override.
	Location string `json:"location"`

	// Data is the data held within a Set.
	Data map[string]Data `json:"data"`
}

// NewSet facilitates the creation of a new Set in a database.
// It's job is to create the necessary directories for the new Set,
// after which it handles updating the its database to account for itself.
//
// Accepts positional arguments: name, uuid, location string.
//
// Note: To save the returned set to disk use the Persist method.
func NewSet(args ...string) (*Set, error) {
	if len(args) == 0 {
		return nil, errors.New("no name provided for new Set.")
	}

	newSet := Set{
		Uuid:     utils.GetUuid(),
		CTime:    time.Now().String(),
		UTime:    time.Now().String(),
		Location: os.Getenv("DEFAULT_DATA_LOCATION"),
	}

	newSet.Name = args[0]

	if len(args) >= 2 {
		newSet.Uuid = args[1]
	}

	if len(args) >= 3 {
		newSet.Location = args[2]
	}

	return &newSet, nil
}

// Read provides access to a Set's meta data.
func (c *Set) Read() (*Set, error) {
	meta := Set{
		Uuid:  c.Uuid,
		Name:  c.Name,
		CTime: c.CTime,
		UTime: c.UTime,
	}
	return &meta, nil
}

// Edit provides a way to edit a Set's meta data.
func (c *Set) Edit() (*Set, error) {
	meta := Set{
		Uuid:  c.Uuid,
		Name:  c.Name,
		CTime: c.CTime,
		UTime: c.UTime,
	}
	return &meta, nil
}

// Delete provides a way to delete a Set from a database.
func (c *Set) Delete() error {
	return nil
}

// Delete provides a way to delete a Set from a database.
func (c *Set) Persist() error {
	return nil
}

// Delete provides a way to delete a Set from a database.
func (c *Set) Load() (*Set, error) {
	return nil, nil
}

// Delete provides a way to delete a Set from a database.
func (c *Set) Insert(data *Data) (*Set, error) {
	return nil, nil
}

func saveDbFile(file *os.File, col *Set) error {
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
