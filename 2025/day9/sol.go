package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2025/day9/input.txt")
	lines := strings.Split(string(dat), "\n")

	// n^2: inspect every pair and calculate the size
	largest := 0
	for i := 0; i < len(lines); i++ {
		for j := i; j < len(lines); j++ {
			largest = max(largest, calcArea(lines[i], lines[j]))
		}
	}
	fmt.Println("Part 1:", largest)
}

func calcArea(from, to string) int {
	fromSplit := strings.Split(from, ",")
	toSplit := strings.Split(to, ",")
	fromX, fromY := getNum(fromSplit[0]), getNum(fromSplit[1])
	toX, toY := getNum(toSplit[0]), getNum(toSplit[1])
	return abs(fromX-toX+1) * abs(fromY-toY+1)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func getNum(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
