package main

import "fmt"

func main() {
	grid := [9][9]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}
	isDone := solve(&grid)
	if isDone {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Printf("%d", grid[i][j])
			}
			fmt.Println()
		}
	}
}

func check(grid [9][9]int, x, y, value int) bool {
	subX := x - x%3
	subY := y - y%3
	for i := subX; i < subX+3; i++ {
		for j := subY; j < subY+3; j++ {
			if grid[i][j] == value {
				//			fmt.Printf("Subgrid: found %d at (%d, %d)\n", value, i, j)
				return false
			}
		}
	}

	for i := 0; i < 9; i++ {
		if grid[i][y] == value && i != x {
			//		fmt.Printf("Row: found %d at (%d, %d)\n", value, i, y)
			return false
		}
		if grid[x][i] == value && i != y {
			//		fmt.Printf("Col: found %d at (%d, %d)\n", value, x, i)
			return false
		}
	}

	return true
}

func solve(grid *[9][9]int) bool {
	x, y := findEmpty(*grid)
	if x == 999 {
		return true
	}
	for val := 1; val <= 9; val++ {
		if check(*grid, x, y, val) {
			grid[x][y] = val
			if solve(grid) {
				return true
			}
			grid[x][y] = 0
		}
	}
	return false
}

func findEmpty(grid [9][9]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return 999, 999
}
