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
		ans += findLargest(line, m, 2)
	}
	fmt.Println("Part 1:", ans)

	ans = 0
	for _, line := range lines {
		largest := getNum(findLargestLeftOf(line, 12))
		ans += largest
		fmt.Println("largest:", largest)
	}
	fmt.Println("Part 2:", ans) // 168_359_209_087_217 is too low
}

func findLargestLeftOf(l string, size int) string {
	if size == 1 {
		return maxstring(l)
	}
	considering := l[0 : len(l)-size+1]
	m := maxstring(considering)
	return m + findLargestLeftOf(l[strings.Index(l, m)+1:], size-1)
}

func maxstring(l string) string {
	m := 0
	for _, r := range l {
		m = max(m, getNum(string(r)))
	}
	return strconv.Itoa(m)
}

func findLargest(line string, m map[string]int, size int) int {
	if m[line] != 0 {
		return m[line]
	}
	if len(line) == size {
		return getNum(line)
	}
	f := string(line[0]) + line[2:]
	first := findLargest(f, m, size)

	second := findLargest(line[1:], m, size)
	third := getNum(line[0:size])
	maximum := max(max(first, second), third)
	m[line] = maximum
	return maximum
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
