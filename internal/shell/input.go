package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var open bool = true

func Read() {
	for open {
		reader := bufio.NewReader((os.Stdin))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		process(text)
		fmt.Println()
	}
}

func process(input string) error {
	command := newCommandNode()
	words := strings.Split(input, " ")

	for _, v := range words {
		if found := determineKeyword(v); found {
			checkExit(v)

			command.cmd = v
			continue
		}

		command.args = append(command.args, v)
	}

	fmt.Print(command)

	return nil
}

func determineKeyword(word string) bool {
	switch word {
	case "exit":
		return true
	}

	return false
}

func checkExit(word string) {
	if word == "exit" {
		open = false
	}
}
