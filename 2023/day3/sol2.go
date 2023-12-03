package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	dat, _ := os.ReadFile("2023/day3/input.txt")
	raw := strings.Split(string(dat), "\n")

	m := make(map[string]int)
	sumParts := 0
	for idx := range raw {
		sumParts += processLine(m, raw, idx)
	}
	fmt.Println("part 2:", sumParts)
}

func processLine(m map[string]int, s []string, row int) int {
	sumParts := 0
	cols := len(s[0])
	startingColumns := make([]int, 0)
	for col := 0; col < len(s[0]); {
		if unicode.IsDigit(rune(s[row][col])) {
			startingColumns = append(startingColumns, col)
			r := regexp.MustCompile("(\\d*)")
			match := r.FindStringSubmatch(s[row][col:cols])[1]
			num, _ := strconv.Atoi(match)
			gear, ok := getGear(s, row, col, col+len(match)-1)
			if ok {
				gearRow, gearCol := gear[0], gear[1]
				key := strconv.Itoa(gearRow) + "." + strconv.Itoa(gearCol)
				if val, ok := m[key]; ok {
					sumParts += num * val
				} else {
					m[key] = num
				}
			}
			col += len(match)
		} else {
			col++
		}
	}
	return sumParts
}

func getGear(s []string, row, colStart, colEnd int) ([]int, bool) {
	// check above, below, sides, diagonal
	if row > 0 {
		for c := colStart; c <= colEnd; c++ {
			if isGear(rune(s[row-1][c])) {
				return []int{row - 1, c}, true
			}
		}
		if colStart > 0 {
			if isGear(rune(s[row-1][colStart-1])) {
				return []int{row - 1, colStart - 1}, true
			}
		}
		if colEnd < len(s[0])-1 {
			if isGear(rune(s[row-1][colEnd+1])) {
				return []int{row - 1, colEnd + 1}, true
			}
		}
	}
	if row < len(s)-1 {
		for c := colStart; c <= colEnd; c++ {
			if isGear(rune(s[row+1][c])) {
				return []int{row + 1, c}, true
			}
		}
		if colStart > 0 {
			if isGear(rune(s[row+1][colStart-1])) {
				return []int{row + 1, colStart - 1}, true
			}
		}
		if colEnd < len(s[0])-1 {
			if isGear(rune(s[row+1][colEnd+1])) {
				return []int{row + 1, colEnd + 1}, true
			}
		}
	}
	if colStart > 0 {
		if isGear(rune(s[row][colStart-1])) {
			return []int{row, colStart - 1}, true
		}
	}
	if colEnd < len(s[0])-1 {
		if isGear(rune(s[row][colEnd+1])) {
			return []int{row, colEnd + 1}, true
		}
	}
	return []int{0, 0}, false
}

func isGear(r rune) bool {
	return r == '*'
}
