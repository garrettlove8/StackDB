package main

import (
	"StackDB/cmd"
	"StackDB/internal/database"
	"StackDB/internal/utils"
	"fmt"
)

var activeDatabase database.Database

func main() {
	utils.GetEnv()

	fmt.Println("stackdb:main")

	cmd.Execute()

	// err := install.Intall()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// system := database.Database{
	// 	Name: "system",
	// }
	// system.Load()

	// fmt.Println("main:system: ", system)
}
