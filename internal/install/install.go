package install

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	version = "0.1"
)

func CheckIntall() error {
	if _, err := os.Stat("./stackdb"); os.IsExist(err) {
		return err
	}

	err := setupDirStructure()
	if err != nil {
		fmt.Println("Unable to setup directory structure: ", err)
	}

	file, err := touchConfigFile()
	if err != nil {
		fmt.Println("Unable to touch config file: ", err)
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
	configFile, err := ioutil.ReadFile(pwd + "/configs/" + version + "/stackdb.yaml")
	if err != nil {
		return err
	}

	_, err = file.Write(configFile)
	if err != nil {
		return err
	}

	return nil
}
