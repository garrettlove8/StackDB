package set

import (
	"time"

	"github.com/google/uuid"
)

// Data is the type used for storing data within collections.
// You can think of Data as being akin to a row in a relational database.
type Data struct {
	// Uuid (Universal Unique Identifier) is the ID for a given Data.
	// This field is internally managed.
	Uuid string `json:"uuid"`

	// CTime (Creation Time) is the time at which the Data was created.
	// This field is internally managed.
	CTime string `json:"cTime"`

	// UTime (Updated Time) is the time at which the Data was created.
	// This field is internally managed.
	UTime string `json:"mTime"`

	// Body is the actual data you are storing in a collection.
	Body map[string][]byte `json:"body"`
}

// NewData is a factory to function that handles the creation on a new data instance.
func NewData() *Data {
	return &Data{
		Uuid:  uuid.New().String(),
		CTime: time.Now().String(),
		UTime: time.Now().String(),
	}
}

// Edit allows you to edit a specific piece of data.
func (d *Data) Edit() (*Data, error) {
	data := Data{
		Uuid:  d.Uuid,
		UTime: time.Now().String(),
		Body:  d.Body,
	}
	return &data, nil
}

// Delete allows you to delete Data from a collection.
func (d *Data) Delete() error {
	return nil
}
