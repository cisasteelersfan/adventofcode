package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2025/day5/small.txt")
	two := strings.Split(string(dat), "\n\n")
	ranges := make([]Range, 0)
	for _, line := range strings.Split(two[0], "\n") {
		s := strings.Split(line, "-")
		ranges = append(ranges, Range{getNum(s[0]), getNum(s[1])})
	}

	fresh := 0
	for _, line := range strings.Split(two[1], "\n") {
		num := getNum(line)
		isFresh := false
		for _, r := range ranges {
			if num >= r.start && num <= r.end {
				isFresh = true
			}
		}
		if isFresh {
			fresh++
		}
	}
	fmt.Println("Part 1:", fresh)

	myrange := Range{ranges[0].start, ranges[0].end}
	for _, r := range ranges {
		if myrange.start > r.start {
			myrange.start = r.start
		}
		if myrange.end < r.end {
			myrange.end = r.end
		}
	}
	fmt.Println("Part 2:", myrange.end-myrange.start+1)
}

type Range struct {
	start, end int
}

func getNum(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}
