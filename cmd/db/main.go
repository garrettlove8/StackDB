package main

import (
	"StackDB/internal/database"
	"StackDB/internal/install"
	"StackDB/internal/utils"
	"fmt"
)

var activeDatabase database.Database

func main() {
	utils.GetEnv()

	err := install.Intall()
	if err != nil {
		fmt.Println(err)
	}

	system := database.Database{
		Name: "system",
	}
	system.Load()

	fmt.Println("main:system: ", system)
}
