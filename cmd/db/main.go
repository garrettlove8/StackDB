package main

import (
	"StackDB/internal/database"
	"StackDB/internal/install"
	"StackDB/internal/utils"
	"fmt"
)

var activeDatabase database.ActiveDatabase

func main() {
	utils.GetEnv()

	err := install.Intall()
	if err != nil {
		fmt.Println(err)
	}

	system := database.ActiveDatabase{
		Name: "system",
	}
	system.Load()

	fmt.Println("main:system: ", system)
}
