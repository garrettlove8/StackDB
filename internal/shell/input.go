package shell

import (
	"StackDB/internal/collection"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

var Open bool = true
var systemCollection *collection.Collection
var activeCollection *collection.Collection

// Start load the stackdb database into memory allowing the user
// to begin using StackDB. After that, it starts the StackDB shell.
func Start() error {
	err := loadSystemDb()
	if err != nil {
		return err
	}

	err = read()
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
		fmt.Println()
	}

	return nil
}

func loadSystemDb() error {
	wantedDb := collection.Collection{
		Name: "stackdb",
	}

	var err error
	if systemCollection, err = wantedDb.Load(); err != nil {
		errMsg := fmt.Sprintf("unable to load database %s", wantedDb.Name)
		return errors.New(errMsg)
	}

	return nil
}

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

	wantedDb := collection.Collection{
		Name: words[1],
	}

	var err error
	activeCollection, err = wantedDb.Load()
	if err != nil {
		newDatabase := collection.NewCollection()
		newDatabase.Name = words[1]
		activeCollection, _ = newDatabase.Create()

		body := make(map[string][]byte)
		body["name"] = []byte(newDatabase.Name)

		newData := collection.NewData()
		newData.CTime = time.Now().String()
		newData.UTime = time.Now().String()
		newData.Body = body

		_, err = systemCollection.Insert(newData)
		if err != nil {
			// TODO: Idealy if there is an error here the process should be undone automatically.

			return fmt.Errorf(`
			database has been created,
			however there was an error adding the new database to the tracking system: %v.
			It is recommended to delete the new database and fix the tracking issue before recreating it.`,
				err)
		}

		err = systemCollection.Persist()
		if err != nil {
			// TODO: Idealy if there is an error here the process should be undone automatically.

			return fmt.Errorf(`
			database has been created,
			however there was an error persisting the new database to the tracking system: %v.
			It is recommended to delete the new database and fix the tracking issue before recreating it.`,
				err)
		}
	}

	return nil
}
