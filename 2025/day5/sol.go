package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2025/day5/input.txt")
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

	// stack blocks!
	starts := make(map[int]int)
	stops := make(map[int]int)
	for _, r := range ranges {
		starts[r.start]++
		stops[r.end]++
	}
	all := make([]int, 0)
	for key := range starts {
		all = append(all, key)
	}
	for key := range stops {
		all = append(all, key)
	}
	sort.Ints(all)
	all = slices.Compact(all)
	depth := 0
	curStart := 0
	total := 0
	for _, key := range all {
		if starts[key] > 0 {
			if depth == 0 {
				curStart = key
			}
			depth += starts[key]
		}
		if stops[key] > 0 {
			depth -= stops[key]
			if depth == 0 {
				total += (key - curStart + 1)
			}
		}
	}
	fmt.Println("Part 2:", total)
}

type Range struct {
	start, end int
}

func getNum(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}
