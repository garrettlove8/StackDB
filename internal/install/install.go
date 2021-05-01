package install

import (
	"io/ioutil"
	"os"
)

// Intall handles the entire install process. This inlcudes
// checking if the database is already installed. If not,
// it kicks off the installation process. However, if it
// was already installed it returns nil allowing for the
// database to start up normally.
func Intall() error {
	if _, err := os.Stat("./stackdb"); os.IsExist(err) {
		return nil
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

	err = os.MkdirAll("./stackdb/data/system", 0777)
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
	file, err := os.Create("./stackdb/config/stackdb.yaml")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func writeInitialConfig(file *os.File) error {
	pwd, _ := os.Getwd()
	configFile, err := ioutil.ReadFile(pwd + "/configs/" + os.Getenv("VERSION") + "/stackdb.yaml")
	if err != nil {
		return err
	}

	_, err = file.Write(configFile)
	if err != nil {
		return err
	}

	return nil
}
