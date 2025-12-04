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
	m := make(map[string]int)
	for _, line := range lines {
		ans += findLargest(line, m)
	}
	fmt.Println("Part 1:", ans)
}

func findLargest(line string, m map[string]int) int {
	if m[line] != 0 {
		return m[line]
	}
	if len(line) == 2 {
		m[line] = getNum(line)
		return getNum(line)
	}
	f := string(line[0]) + line[2:]
	first := findLargest(f, m)
	m[f] = first

	second := findLargest(line[1:], m)
	m[line[1:]] = second
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
