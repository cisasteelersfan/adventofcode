package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	dat, _ := os.ReadFile("2023/day1/input.txt")
	raw := strings.Split(string(dat), "\n")

	total := 0

	for _, line := range raw {
		total += computeFirstLast(line)
	}
	fmt.Println(total)
}

func computeFirstLast(s string) int {
	first := getFirst(s)
	last := getLast(s)
	return first*10 + last
}

func getFirst(s string) int {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}
	panic("no digits found")
}

func getLast(s string) int {
	for idx := range s {
		char := rune(s[len(s)-idx-1])
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}
	panic("no digits found")
}
