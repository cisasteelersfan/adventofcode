package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2025/day3/input.txt")
	lines := strings.Split(string(dat), "\n")

	ans := 0
	for _, line := range lines {
		ans += findLargest(line)
	}
	fmt.Println("Part 1:", ans)
}

func findLargest(line string) int {
	if len(line) == 2 {
		return getNum(line)
	}
	first := findLargest(string(line[0]) + line[2:])
	second := findLargest(line[1:])
	third := getNum(line[0:2])
	return max(max(first, second), third)
}

func getNum(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
