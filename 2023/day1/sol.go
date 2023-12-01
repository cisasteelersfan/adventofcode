package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	dat, _ := os.ReadFile("2023/day1/input.txt")
	raw := strings.Split(string(dat), "\n")

	total := 0

	for _, line := range raw {
		total += computeFirstLast2(line)
	}
	fmt.Println("part 2:", total)
}

func computeFirstLast2(s string) int {
	first := getFirst2(s)
	last := getLast2(s)
	ans := first*10 + last
	// fmt.Println("debug: ", s, ans)
	return ans
}

func getFirst2(s string) int {
	for index, char := range s {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
		// otherwise look ahead 2-4 characters
		if getTextNum(s, index) != 0 {
			return getTextNum(s, index)
		}
	}
	panic("no first digit found in string: " + s)
}

func getTextNum(s string, i int) int {
	s = s + "    "
	num, ok := digits[s[i:i+3]]
	if ok {
		return num
	}
	num, ok = digits[s[i:i+4]]
	if ok {
		return num
	}
	return digits[s[i:i+5]]
}

func getLast2(s string) int {
	for idx := range s {
		index := len(s) - idx - 1
		char := rune(s[index])
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
		if getTextNum(s, index) != 0 {
			return getTextNum(s, index)
		}
	}
	panic("no last digit found in string: " + s)
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
