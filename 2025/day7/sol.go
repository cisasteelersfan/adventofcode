package main

import (
	"fmt"
	"os"
	"strconv"
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

	g = make([][]string, 0)
	for _, line := range lines {
		g = append(g, make([]string, 0))
		for _, char := range line {
			g[len(g)-1] = append(g[len(g)-1], string(char))
		}
	}
	// approach: count the number of lines hitting the bottom
	for row := 1; row < len(g); row++ {
		for col := 0; col < len(g[0]); col++ {
			char := g[row][col]
			if char == "." {
				// look up
				if g[row-1][col] == "S" {
					g[row][col] = "1"
				} else if isNum(g[row-1][col]) {
					g[row][col] = g[row-1][col]
				}
			}
		}
		for col := 0; col < len(g[0]); col++ {
			char := g[row][col]
			if char == "^" {
				if isNum(g[row-1][col]) {
					g[row][col-1] = strconv.Itoa(getNumOrZero(g[row][col-1]) + getNum(g[row-1][col]))
					g[row][col+1] = strconv.Itoa(getNumOrZero(g[row][col+1]) + getNum(g[row-1][col]))
				}
			}
		}
	}
	ans := 0
	for _, s := range g[len(g)-1] {
		ans += getNumOrZero(s)
	}
	fmt.Println("Part 2:", ans)
}

func isNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func getNum(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func getNumOrZero(s string) int {
	if isNum(s) {
		return getNum(s)
	}
	return 0
}

func printGraph(g [][]string) {
	for line := range g {
		fmt.Println(g[line])
	}
}
