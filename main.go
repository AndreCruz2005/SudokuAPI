package main

import (
	"SudokuAPI/Generator"
	"fmt"
)

func main() {
	g := Generator.GenerateSudoku(80)

	for i := 0; i < 9; i++ {

		fmt.Println(g[i])

	}
}
