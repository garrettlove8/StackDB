package main

import (
	"StackDB/internal/database"
	"StackDB/internal/install"
	"StackDB/internal/utils"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("StackDB says hello!")
	utils.GetEnv()

	systemDb := database.Database{
		Id:   uuid.New().String(),
		Name: "system",
		Type: "keyValue",
	}
	databasesCol := database.Collection{
		Id:   uuid.New().String(),
		Name: "databases",
	}
	install.Intall(&systemDb, &databasesCol)
}
