package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day4/input.txt")
	raw := strings.Split(string(dat), "\n")

	total := 0
	for _, line := range raw {
		if overlaps(line) {
			total += 1
		}
	}
	fmt.Println("part 1:", total)

	total = 0
	for _, line := range raw {
		if overlapAtAll(line) {
			total += 1
		}
	}
	fmt.Println("part 2:", total)
}

func overlapAtAll(l string) bool {
	s := strings.Split(l, ",")
	firstStart, firstEnd := parseStartEnd(s[0])
	secStart, secEnd := parseStartEnd(s[1])
	return (firstStart <= secStart && secStart <= firstEnd) || (firstStart <= secEnd && secEnd <= firstEnd) ||
		(secStart <= firstStart && firstStart <= secEnd) || (secStart <= firstEnd && firstEnd <= secEnd)
}

func overlaps(l string) bool {
	s := strings.Split(l, ",")
	firstStart, firstEnd := parseStartEnd(s[0])
	secStart, secEnd := parseStartEnd(s[1])
	return (firstStart <= secStart && firstEnd >= secEnd) || (secStart <= firstStart && secEnd >= firstEnd)
}

func parseStartEnd(l string) (int, int) {
	s := strings.Split(l, "-")
	start, _ := strconv.Atoi(s[0])
	end, _ := strconv.Atoi(s[1])
	return start, end
}
