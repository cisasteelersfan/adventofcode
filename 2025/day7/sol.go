package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2025/day7/input.txt")
	lines := strings.Split(string(dat), "\n")
	g := make([][]string, 0)
	for _, line := range lines {
		g = append(g, make([]string, 0))
		for _, char := range line {
			g[len(g)-1] = append(g[len(g)-1], string(char))
		}
	}
	splits := 0
	for row := 1; row < len(g); row++ {
		for col := 0; col < len(g[0]); col++ {
			char := g[row][col]
			if char == "." {
				// look up
				if g[row-1][col] == "S" || g[row-1][col] == "|" {
					g[row][col] = "|"
				}
			} else if char == "^" {
				if g[row-1][col] == "|" {
					g[row][col-1] = "|"
					g[row][col+1] = "|"
					splits++
				}
			}
		}
	}
	fmt.Println("Part 1:", splits)
}
