package database

import "encoding/json"

type ActiveDatabase struct {
	Id          string
	Name        string
	Type        string
	ColStrings  []string
	Collections []Collection
}

func (ad *ActiveDatabase) Load() error {
	db, err := loadDatabase(ad.Name)
	if err != nil {
		return err
	}

	ad.Id = db.Id
	ad.Type = db.Type

	for _, v := range db.Collections {
		data := readColFile(db.Name, v)
		col, _ := encodeColFile(data)
		ad.Collections = append(ad.Collections, *col)
	}

	return nil
}

func loadDatabase(dbName string) (*Database, error) {
	var database Database
	dataBytes := readDbFile(dbName)

	err := json.Unmarshal(dataBytes, &database)
	if err != nil {
		return nil, err
	}

	return &database, nil
}
