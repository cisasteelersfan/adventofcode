package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2025/day4/input.txt")
	lines := strings.Split(string(dat), "\n")
	graph := make(map[Node]string)
	for row, line := range lines {
		for col, char := range line {
			graph[Node{row: row, col: col}] = string(char)
		}
	}
	forklifts := 0
	for row, line := range lines {
		for col := range line {
			if graph[Node{row: row, col: col}] != "@" {
				continue
			}
			if countRolls(graph, row, col) < 4 {
				forklifts++
			}
		}
	}
	fmt.Println("Part 1:", forklifts)

	rollsRemoved := 0
	changed := true
	for changed {
		changed = false

		forklifts := 0
		for row, line := range lines {
			for col := range line {
				if graph[Node{row: row, col: col}] != "@" {
					continue
				}
				if countRolls(graph, row, col) < 4 {
					forklifts++
					changed = true
					graph[Node{row: row, col: col}] = "."
				}
			}
		}
		rollsRemoved += forklifts
	}
	fmt.Println("Part 2:", rollsRemoved)
}

func countRolls(graph map[Node]string, row, col int) int {
	rolls := 0
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}
			if graph[Node{row: row + r, col: col + c}] == "@" {
				rolls++
			}
		}
	}
	return rolls
}

type Node struct {
	row, col int
}
