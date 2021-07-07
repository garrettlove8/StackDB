package set

import (
	"encoding/json"
	"errors"
	"os"
)

// Sets are the same as those in many NoSQL database and are akin to tables in relational databases.
// They privde an easy and logical way to separate data with a database
type Set struct {
	// Header contains metadata for a set.
	Header Header

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
func NewSet(name, version, pk, id string) (*Set, error) {
	if name == "" {
		// TODO: Eventially, this should also log to the error log
		return nil, errors.New("No name provided")
	}

	newHeader := newHeader(name, version, pk, id)

	newSet := Set{
		Header: *newHeader,
	}

	return &newSet, nil
}

// Read provides access to a Set's meta data.
func (c *Set) Read() (*Set, error) {
	meta := Set{}
	return &meta, nil
}

// Edit provides a way to edit a Set's meta data.
func (c *Set) Edit() (*Set, error) {
	meta := Set{}
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
