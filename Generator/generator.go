package Generator

import "math/rand/v2"

func notInBox(grid *[9][9]int, row int, col int, num int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[row+i][col+j] == num {
				return false
			}
		}
	}
	return true
}

func notInRow(grid *[9][9]int, row int, num int) bool {
	for col := 0; col < 9; col++ {
		if grid[row][col] == num {
			return false
		}
	}
	return true
}

func notInCol(grid *[9][9]int, col int, num int) bool {
	for row := 0; row < 9; row++ {
		if grid[row][col] == num {
			return false
		}
	}
	return true
}

func canBePlaced(grid *[9][9]int, row int, col int, num int) bool {
	return notInRow(grid, row, num) &&
		notInCol(grid, col, num) &&
		notInBox(grid, row-row%3, col-col%3, num)
}

func fillBox(grid *[9][9]int, row int, col int) {
	num := rand.IntN(10)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for !notInBox(grid, row, col, num) {
				num = rand.IntN(10)
			}
			grid[row+i][col+j] = num
		}
	}
}

func fillDiagonal(grid *[9][9]int) {
	for i := 0; i < 9; i += 3 {
		fillBox(grid, i, i)
	}
}

func fillOthers(grid *[9][9]int, row int, col int) bool {
	if row == 9 {
		return true
	}

	if col == 9 {
		return fillOthers(grid, row+1, 0)
	}

	if grid[row][col] != 0 {
		return fillOthers(grid, row, col+1)
	}

	for num := 1; num <= 9; num++ {
		if canBePlaced(grid, row, col, num) {
			grid[row][col] = num
			if fillOthers(grid, row, col+1) {
				return true
			}
			grid[row][col] = 0
		}
	}

	return false
}

func hideNumbers(grid *[9][9]int, hints int) {
	var toRemove = 81 - hints
	for toRemove > 0 {
		row := rand.IntN(9)
		col := rand.IntN(9)
		if grid[row][col] != 0 {
			grid[row][col] = 0
			toRemove--
		}

	}
}

// OrganizeIntoBoxes Changes the 2d array where each array represents a row to one where each array is a box
func OrganizeIntoBoxes(grid *[9][9]int) {
	var boxes [9][9]int

	// Iterates over all the boxes
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {

			// Creates each box
			var box [9]int

			// Iterates over all cells in the box
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					// Calculate the row and column in the original grid
					row := boxRow*3 + i
					col := boxCol*3 + j
					cell := grid[row][col]
					box[i*3+j] = cell
				}
			}
			boxes[boxRow*3+boxCol] = box
		}
	}
}

func GenerateSudoku(hints int) [9][9]int {
	var grid [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			grid[i][j] = 0
		}
	}

	fillDiagonal(&grid)
	fillOthers(&grid, 0, 0)
	hideNumbers(&grid, hints)

	return grid
}
