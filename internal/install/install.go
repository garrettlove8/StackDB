package install

import (
	"fmt"
	"os"
)

const (
	version = 0.1
)

func CheckIntall() {
	// Check if ".stackdb" directory exists, if not do your thing
	if _, err := os.Stat("./stackdb"); os.IsNotExist(err) {
		fmt.Println("The directory DOES NOT exist")
		err = setupDirStructure()
		if err != nil {
			fmt.Println("Unable to setup directory structure: ", err)
		}
	}
}

func setupDirStructure() error {
	err := os.MkdirAll("./stackdb/logs/transaction", 0777)
	if err != nil {
		fmt.Println("Unable to setup directory logs/transaction: ", err)
	}

	err = os.MkdirAll("./stackdb/logs/debug", 0777)
	if err != nil {
		fmt.Println("Unable to setup directory logs/debug: ", err)
	}

	err = os.MkdirAll("./stackdb/logs/stats", 0777)
	if err != nil {
		fmt.Println("Unable to setup directory logs/stats: ", err)
	}

	err = os.MkdirAll("./stackdb/data/system", 0777)
	if err != nil {
		fmt.Println("Unable to setup directory data/system: ", err)
	}

	// err = os.MkdirAll("./stackdb/config/stac", 0777)

	return err
}
