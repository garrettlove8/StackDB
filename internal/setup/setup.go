package setup

import (
	"StackDB/internal/collections"
	"fmt"
	"io/ioutil"
	"os"
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

	configBytes, err := getConfigContent()
	if err != nil {
		return err
	}

	err = setupDirStructure()
	if err != nil {
		return err
	}

	file, err := touchConfigFile()
	if err != nil {
		return err
	}

	err = writeInitialConfig(file, configBytes)
	if err != nil {
		return err
	}

	systemSets, err := collections.NewSet("sets")
	if err != nil {
		return err
	}

	err = systemSets.Persist()
	if err != nil {
		return err
	}

	return nil
}

// CheckSetup checks to see if the StackDB setup process has been
// previously run.
func CheckSetup() bool {
	if _, err := os.Stat("~/sdb"); !os.IsNotExist(err) {
		return true
	}

	return false
}

func setupDirStructure() error {
	homepath, _ := os.UserHomeDir()
	err := os.Chdir(homepath)
	if err != nil {
		return err
	}

	err = os.MkdirAll("./sdb/logs/transactions", 0777)
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

func getConfigContent() ([]byte, error) {
	pwd, _ := os.Getwd()
	configBytes, err := ioutil.ReadFile(pwd + "/configs/" + os.Getenv("VERSION") + "/stackdb.json")
	if err != nil {
		return nil, err
	}

	return configBytes, nil
}

func writeInitialConfig(file *os.File, configBytes []byte) error {
	_, err := file.Write(configBytes)
	if err != nil {
		return err
	}

	return nil
}
