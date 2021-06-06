package main

import (
	"StackDB/internal/set"
	"StackDB/internal/utils"
	"fmt"
)

// var systemDatabase database.Database
// var activeDatabase database.Database

func main() {
	utils.GetEnv()

	newSet, err := set.NewSet()
	if err != nil {
		fmt.Println("main:err: ", err)
	}

	fmt.Println("main:success:newSet: ", newSet)

	// cmd.Execute()
}
