package main

import (
	"StackDB/internal/install"
	"fmt"
)

func main() {
	fmt.Println("StackDB says hello!")
	install.CheckIntall()
}
