package setup

import (
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
	isSetup := CheckSetup()
	if isSetup {
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

	// systemCollection, err := collections.NewCollection("collectiions")
	// if err != nil {
	// 	return err
	// }

	// err = systemCollection.Persist()
	// if err != nil {
	// 	return err
	// }

	fmt.Println("Setup process complete. You can now use StackDB")

	return nil
}

// CheckSetup checks to see if the StackDB setup process has been
// previously run.
func CheckSetup() bool {
	homepath, _ := os.UserHomeDir()
	err := os.Chdir(homepath)
	if err != nil {
		fmt.Println("1. Unable to check setup status, see logs") // TODO: Logs this to DB logs
	}

	_, err = os.Getwd()
	if err != nil {
		fmt.Println("2. Unable to check setup status, see logs") // TODO: Logs this to DB logs
	}

	_, err = os.Stat("./sdb")

	return !os.IsNotExist(err)
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
	// TODO: Change path here
	configBytes, err := ioutil.ReadFile("/Users/garrettlove/Development/StackDB/configs/" + os.Getenv("VERSION") + "/stackdb.json")
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
