package set

import (
	"errors"

	"github.com/google/uuid"
)

// Header hold all metadata for a given set.
type Header struct {
	// Id is the globally unique identifier for a set.
	// It can be configured if you wish to use your own
	// id system. By default, it uses a V4 UUID.
	Id string

	// Name is the name of the set. There is no default
	// and must be provided. If no name is provided a
	// "No name provided" error will be thrown.
	Name string

	// PrimaryKey is the value that will be used as
	// the key in the data map. By default, the record
	// id is used, however you can configure this if
	// you'll be primarily searching on a different
	// property. At the foundational level this is
	// how an index would be implemented.
	PrimaryKey string
}

// newHeader is a factory function for creating a set header.
// It uses the positional arguments for: Name, Key, and Uuid
func newHeader(name, pk, id string) (*Header, error) {
	if name == "" {
		// TODO: Eventially, this should also log to the error log
		return nil, errors.New("No name provided")
	}

	newHeader := &Header{}

	newHeader.Name = name

	// Allows for setting a custom PrimaryKey when the set is being
	// created and stored. If no custom key is provided the record
	// Id will be used.
	if pk != "" {
		// TODO: Since this is dependant on a specific property being present
		// in a data record we'll need to ensure that this property is present
		// when new data records are being added
		newHeader.PrimaryKey = pk
	} else {
		newHeader.PrimaryKey = "id"
	}

	if id != "" {
		newHeader.Id = id
	} else {
		newUuid, _ := uuid.NewRandom()
		newHeader.Id = newUuid.String()
	}

	return newHeader, nil
}
