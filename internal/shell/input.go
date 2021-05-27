package shell

import (
	"StackDB/internal/database"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

var Open bool = true
var systemDatabase *database.Database
var activeDatabase *database.Database

func Start() error {
	wantedDb := database.Database{
		Name: "stackdb",
	}

	var err error
	if systemDatabase, err = wantedDb.Load(); err != nil {
		errMsg := fmt.Sprintf("unable to load database %s", wantedDb.Name)
		return errors.New(errMsg)
	}

	return nil
}

func Read() error {
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

	fmt.Println("user would like to use a database")

	wantedDb := database.Database{
		Name: words[1],
	}

	var err error
	activeDatabase, err = wantedDb.Load()
	if err != nil {
		newDatabase := database.NewDatabase()
		newDatabase.Name = words[1]
		newDatabase.Type = words[2]
		activeDatabase, _ = newDatabase.Create()

		body := make(map[string][]byte)
		body["name"] = []byte(newDatabase.Name)

		newData := database.NewData()
		newData.CTime = time.Now().String()
		newData.MTime = time.Now().String()
		newData.Body = body

		_, err = systemDatabase.Insert("databases", newData)
		if err != nil {
			// TODO: Idealy if there is an error here the process should be undone automatically.

			return fmt.Errorf(`
			database has been created,
			however there was an error adding the new database to the tracking system: %v.
			It is recommended to delete the new database and fix the tracking issue before recreating it.`,
				err)
		}

		err = systemDatabase.Persist()
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
