package main

import (
	"SudokuAPI/Server"
	"fmt"
)

func main() {
	fmt.Println("Server running")
	Server.Serve()
}
