package main

import (
	"StackDB/cmd"
	"StackDB/internal/database"
	"StackDB/internal/utils"
	"fmt"
	"os"
)

var systemDatabase database.Database
var activeDatabase database.Database

func main() {
	utils.GetEnv()

	err := cmd.Execute(&activeDatabase, &systemDatabase)
	if err != nil {
		fmt.Println("main:Execute:error: ", err)
		os.Exit(1)
	}
}
