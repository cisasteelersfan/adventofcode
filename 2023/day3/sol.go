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

	sumParts := 0
	for idx := range raw {
		sumParts += processLine(raw, idx)
	}
	fmt.Println("part 1:", sumParts)
}

func processLine(s []string, row int) int {
	cols := len(s[0])
	sumParts := 0
	startingColumns := make([]int, 0)
	for col := 0; col < len(s[0]); {
		if unicode.IsDigit(rune(s[row][col])) {
			startingColumns = append(startingColumns, col)
			r := regexp.MustCompile("(\\d*)")
			match := r.FindStringSubmatch(s[row][col:cols])[1]
			num, _ := strconv.Atoi(match)
			if isValid(s, row, col, col+len(match)-1) {
				sumParts += num
			}
			col += len(match)
		} else {
			col++
		}
	}
	return sumParts
}

func isValid(s []string, row, colStart, colEnd int) bool {
	// check above, below, sides, diagonal
	if row > 0 {
		for c := colStart; c <= colEnd; c++ {
			if isSymbol(rune(s[row-1][c])) {
				return true
			}
		}
		if colStart > 0 {
			if isSymbol(rune(s[row-1][colStart-1])) {
				return true
			}
		}
		if colEnd < len(s[0])-1 {
			if isSymbol(rune(s[row-1][colEnd+1])) {
				return true
			}
		}
	}
	if row < len(s)-1 {
		for c := colStart; c <= colEnd; c++ {
			if isSymbol(rune(s[row+1][c])) {
				return true
			}
		}
		if colStart > 0 {
			if isSymbol(rune(s[row+1][colStart-1])) {
				return true
			}
		}
		if colEnd < len(s[0])-1 {
			if isSymbol(rune(s[row+1][colEnd+1])) {
				return true
			}
		}
	}
	if colStart > 0 {
		if isSymbol(rune(s[row][colStart-1])) {
			return true
		}
	}
	if colEnd < len(s[0])-1 {
		if isSymbol(rune(s[row][colEnd+1])) {
			return true
		}
	}
	return false
}

func isSymbol(r rune) bool {
	return r != '.' && !unicode.IsDigit(r)
}
