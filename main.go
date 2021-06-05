package main

import (
	"StackDB/internal/set"
	"StackDB/internal/utils"
)

// var systemDatabase database.Database
// var activeDatabase database.Database

func main() {
	utils.GetEnv()

	newSet := set.NewSet()
	newSet.Create("first", "heyUUID")

	// cmd.Execute()
}
