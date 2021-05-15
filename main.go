package main

import (
	"StackDB/cmd"
	"StackDB/internal/database"
	"StackDB/internal/utils"
)

var systemDatabase database.Database
var activeDatabase database.Database

func main() {
	utils.GetEnv()

	cmd.Execute(&activeDatabase, &systemDatabase)
}
