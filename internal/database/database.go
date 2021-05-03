package database

import (
	"fmt"
	"os"
)

type Database struct {
	Id   string
	Name string
	Type string
}

func (db *Database) Create() error {
	fmt.Println("database:Create:db: ", db)

	err := createDatabaseDir(db.Name)
	if err != nil {
		return err
	}

	err = createDbFile(*db)
	if err != nil {
		return err
	}

	err = createCollectionsDir(db.Name)
	if err != nil {
		return err
	}

	return nil
}

func createDatabaseDir(dbName string) error {
	err := os.MkdirAll("./stackdb/data/"+dbName, 0777)
	if err != nil {
		return err
	}

	return nil
}

func createCollectionsDir(dbName string) error {
	err := os.MkdirAll("./stackdb/data/"+dbName+"/collections", 0777)
	if err != nil {
		return err
	}

	return nil
}

func createDbFile(db Database) error {
	file, err := os.Create("./stackdb/data/" + db.Name + "/database.json")
	if err != nil {
		return err
	}

	content := fmt.Sprintf("{\n\t\"uuid\": \"%s\", \n\t\"name\": \"%s\", \n\t\"type\": \"%s\"\n}", db.Id, db.Name, db.Type)
	contentBytes := []byte(content)

	_, err = file.Write(contentBytes)
	if err != nil {
		return err
	}

	return nil
}
