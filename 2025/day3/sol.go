package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2025/day3/small.txt")
	lines := strings.Split(string(dat), "\n")

	ans := 0
	m := make(map[string]int)
	for _, line := range lines {
		ans += findLargest(line, m, 2)
	}
	fmt.Println("Part 1:", ans)

	ans = 0
	for _, line := range lines {
		n := make(map[string]int)
		largest := findLargest(line, n, 12)
		ans += largest
		fmt.Println("largest:", largest)
	}
	fmt.Println("Part 2:", ans) // 168_359_209_087_217 is too low
}

func findLargest(line string, m map[string]int, size int) int {
	if m[line] != 0 {
		return m[line]
	}
	if len(line) == size {
		m[line] = getNum(line)
		return getNum(line)
	}
	f := string(line[0]) + line[2:]
	first := findLargest(f, m, size)
	m[f] = first

	second := findLargest(line[1:], m, size)
	m[line[1:]] = second
	third := getNum(line[0:size])
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
