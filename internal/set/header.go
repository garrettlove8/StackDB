package set

import (
	"github.com/google/uuid"
)

// Header hold all metadata for a given set.
type Header struct {
	// Id is the globally unique identifier for a set.
	// It can be configured if you wish to use your own
	// id system. By default, it uses a V4 UUID.
	Id string `json:"id"`

	// Version is the version of the set.
	Version string `json:"version"`

	// Location is the directory path to where the set is stored on disk.
	// By default, this is managed internally by StackDB, however, if you
	// are building on top of StackDB this may be helpful to override.
	Location string `json:"location"`

	// Name is the name of the set. There is no default
	// and must be provided. If no name is provided a
	// "No name provided" error will be thrown.
	Name string `json:"name"`

	// PrimaryKey is the value that will be used as
	// the key in the data map. By default, the record
	// id is used, however you can configure this if
	// you'll be primarily searching on a different
	// property. At the foundational level this is
	// how an index would be implemented.
	PrimaryKey string `json:"primaryKey"`
}

// newHeader is a factory function for creating a set header.
// It uses the positional arguments for: Name, Key, and Uuid
func newHeader(name, version, pk, id string) *Header {
	newHeader := &Header{}

	newHeader.Name = name

	if version != "" {
		newHeader.Version = version
	} else {
		// TODO: Hardcoded to 1 for now, should be getting value from
		// config file in the future.
		newHeader.Version = "1"
	}

	// Allows for setting a custom PrimaryKey when the set is being
	// created and stored. If pk is not provided the record
	// Id will be used.
	if pk != "" {
		// TODO: Since this is dependant on a specific property being present
		// in a data record we'll need to ensure that this property is present
		// when new data records are being added
		newHeader.PrimaryKey = pk
	} else if pk == "defualt" {
		newHeader.PrimaryKey = "id"
	}

	if id != "" {
		newHeader.Id = id
	} else if id == "defualt" {
		newUuid, _ := uuid.NewRandom()
		newHeader.Id = newUuid.String()
	}

	return newHeader
}
