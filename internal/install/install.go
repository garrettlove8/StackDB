package install

import (
	"StackDB/internal/database"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

// Intall handles the entire install process. This inlcudes
// checking if the database is already installed. If not,
// it kicks off the installation process. However, if it
// was already installed it returns nil allowing for the
// database to start up normally.
func Intall() error {
	if _, err := os.Stat("./stackdb"); !os.IsNotExist(err) {
		return nil
	}

	db := database.Database{
		Uuid: uuid.New().String(),
		Name: "system",
		Type: "keyValue",
	}
	col := database.Collection{
		Uuid: uuid.New().String(),
		Name: "databases",
	}

	err := setupDirStructure()
	if err != nil {
		return err
	}

	file, err := touchConfigFile()
	if err != nil {
		return err
	}

	err = writeInitialConfig(file)
	if err != nil {
		return err
	}

	newDb, err := db.Create()
	if err != nil {
		return err
	}

	col.Create(newDb.Name)

	return nil
}

func setupDirStructure() error {
	err := os.MkdirAll("./stackdb/logs/transaction", 0777)
	if err != nil {
		return err
	}

	err = os.MkdirAll("./stackdb/logs/debug", 0777)
	if err != nil {
		return err
	}

	err = os.MkdirAll("./stackdb/logs/stats", 0777)
	if err != nil {
		return err
	}

	err = os.MkdirAll("./stackdb/data", 0777)
	if err != nil {
		return err
	}

	err = os.MkdirAll("./stackdb/config", 0777)
	if err != nil {
		return err
	}

	return nil
}

func touchConfigFile() (*os.File, error) {
	file, err := os.Create("./stackdb/config/stackdb.json")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func writeInitialConfig(file *os.File) error {
	pwd, _ := os.Getwd()
	configFile, err := ioutil.ReadFile(pwd + "/configs/" + os.Getenv("VERSION") + "/stackdb.json")
	if err != nil {
		return err
	}

	_, err = file.Write(configFile)
	if err != nil {
		return err
	}

	return nil
}
