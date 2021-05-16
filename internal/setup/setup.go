package setup

import (
	"StackDB/internal/database"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

// Setup handles the entire install process. This inlcudes
// checking if the database is already installed. If not,
// it kicks off the installation process. However, if it
// was already installed it returns nil allowing for the
// database to start up normally.
func Setup() error {
	if isSetup := CheckSetup(); isSetup {
		fmt.Println("StackDB has already been setup")
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

// CheckSetup checks to see if the StackDB setup process has been
// previously run.
func CheckSetup() bool {
	if _, err := os.Stat("./sdb"); !os.IsNotExist(err) {
		return true
	}

	return false
}

func setupDirStructure() error {
	err := os.MkdirAll("./sdb/logs/transaction", 0777)
	if err != nil {
		return err
	}

	err = os.MkdirAll("./sdb/logs/debug", 0777)
	if err != nil {
		return err
	}

	err = os.MkdirAll("./sdb/logs/stats", 0777)
	if err != nil {
		return err
	}

	err = os.MkdirAll("./sdb/data", 0777)
	if err != nil {
		return err
	}

	err = os.MkdirAll("./sdb/config", 0777)
	if err != nil {
		return err
	}

	return nil
}

func touchConfigFile() (*os.File, error) {
	file, err := os.Create("./sdb/config/stackdb.json")
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
