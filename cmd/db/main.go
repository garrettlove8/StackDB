package main

import (
	"StackDB/internal/install"
	"StackDB/internal/utils"
	"fmt"
)

func main() {
	fmt.Println("StackDB says hello!")
	utils.GetEnv()
	install.Intall()
}
