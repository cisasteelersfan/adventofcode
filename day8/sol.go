package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day8/input.txt")
	raw := strings.Split(string(dat), "\n")
	// approach: copy to an empty grid of booleans. Go from each direction: left, right, up down,
	// keeping track of the highest encountered so far. Mark visible if greater than.
	trees := parseTrees(raw)
	visible := findVisible(trees)
	fmt.Println("part 1:", getSum(visible))
}

func getSum(grid [][]bool) int {
	rows := len(grid)
	total := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < rows; c++ {
			if grid[r][c] {
				total++
			}
		}
	}
	return total
}

func findVisible(trees [][]int) [][]bool {
	rows := len(trees)
	grid := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]bool, rows)
	}
	// from left
	for row := 0; row < rows; row++ {
		biggestSoFar := -1
		for col := 0; col < rows; col++ {
			if trees[row][col] > biggestSoFar {
				grid[row][col] = true
				biggestSoFar = trees[row][col]
			}
		}
	}
	// from right
	for row := 0; row < rows; row++ {
		biggestSoFar := -1
		for col := rows - 1; col >= 0; col-- {
			if trees[row][col] > biggestSoFar {
				grid[row][col] = true
				biggestSoFar = trees[row][col]
			}
		}
	}
	// from top
	for col := 0; col < rows; col++ {
		biggestSoFar := -1
		for row := 0; row < rows; row++ {
			if trees[row][col] > biggestSoFar {
				grid[row][col] = true
				biggestSoFar = trees[row][col]
			}
		}
	}
	// from bottom
	for col := 0; col < rows; col++ {
		biggestSoFar := -1
		for row := rows - 1; row >= 0; row-- {
			if trees[row][col] > biggestSoFar {
				grid[row][col] = true
				biggestSoFar = trees[row][col]
			}
		}
	}
	return grid
}

func parseTrees(r []string) [][]int {
	rows := len(r)
	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]int, rows)
		for j := 0; j < rows; j++ {
			grid[i][j] = getInt(r[i][j])
		}
	}
	return grid
}

func getInt(b byte) int {
	num, _ := strconv.Atoi(string(b))
	return num
}
