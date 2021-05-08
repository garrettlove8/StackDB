package storage

import "StackDB/internal/database"

type ActiveDatabase struct {
	Database    database.Database
	Collections []database.Collection
}

type Loader interface {
	LoadDatabase() (database.Database, error)
	LoadCollections() ([]database.Collection, error)
}

func (ad *ActiveDatabase) Load(db Loader, col Loader) error {
	resDatabase, err := db.LoadDatabase()
	resCollections, err := col.LoadCollections()
	if err != nil {
		return err
	}

	ad.Database = resDatabase
	ad.Collections = resCollections

	return nil
}
