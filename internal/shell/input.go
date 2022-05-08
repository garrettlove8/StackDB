package shell

import (
	"StackDB/internal/collections"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Open bool = true
var systemCollection *collections.Collection
var activeCollection *collections.Collection

// Start load the stackdb database into memory allowing the user
// to begin using StackDB. After that, it starts the StackDB shell.
func Start() error {
	// err := loadSystemDb()
	// if err != nil {
	// 	return err
	// }

	err := read()
	if err != nil {
		return err
	}

	return nil
}

func read() error {
	for Open {
		reader := bufio.NewReader((os.Stdin))
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		text = strings.TrimSuffix(text, "\n")
		process(text)
	}

	return nil
}

// func loadSystemDb() error {
// 	wantedCollection := collections.Collection{
// 		Name: "stackdb",
// 	}

// 	var err error
// 	if systemCollection, err = wantedCollection.Load(); err != nil {
// 		errMsg := fmt.Sprintf("unable to load database %s", wantedCollection.Name)
// 		return errors.New(errMsg)
// 	}

// 	return nil
// }

func process(input string) error {
	command := newCommandNode()

	words := strings.Split(input, " ")
	handleExit(words[0])
	handleUse(words)

	// Reset words to split input using . as the separator
	words = strings.Split(input, ".")

	for _, v := range words {
		if found := determineKeyword(v); found {
			command.cmd = v
			continue
		} else {
			command.args = append(command.args, v)
		}
	}

	fmt.Print(command)

	return nil
}

func determineKeyword(word string) bool {
	// TODO: Would be more efficient to use a map
	switch word {
	case "db":
		return true
	}

	return false
}

func handleExit(word string) {
	if word == "exit" {
		Open = false
	}
}

func handleUse(words []string) error {
	if words[0] != "use" {
		return nil
	}

	wantedCollection := collections.Collection{
		Name: words[1],
	}

	var err error
	_, _, err = wantedCollection.Load()
	if err != nil {
		newCollection, _ := collections.NewCollection(words[1])

		body := make(map[string][]byte)
		body["name"] = []byte(newCollection.Name)

		// newData := collections.NewData()
		// newData.Body = body
		homepath, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Unable to create collection: ", err)
			return err
		}

		file, err := os.Create(homepath + "/sdb/data/" + newCollection.Name + ".json")
		if err != nil {
			fmt.Println("handUse:error: ", err)
			return err
		}

		err = newCollection.Persist(file)
		if err != nil {
			fmt.Println("handUse:error: ", err)
			return err
		}

		return err

		// _, err = systemCollection.Insert(newData)
		// if err != nil {
		// 	// TODO: Idealy if there is an error here the process should be undone automatically.

		// 	return fmt.Errorf(`
		// 	database has been created,
		// 	however there was an error adding the new database to the tracking system: %v.
		// 	It is recommended to delete the new database and fix the tracking issue before recreating it`,
		// 		err)
		// }

		// err = systemCollection.Persist()
		// if err != nil {
		// 	// TODO: Idealy if there is an error here the process should be undone automatically.

		// 	return fmt.Errorf(`
		// 	database has been created,
		// 	however there was an error persisting the new database to the tracking system: %v.
		// 	It is recommended to delete the new database and fix the tracking issue before recreating it`,
		// 		err)
		// }
	}

	return nil
}
