package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day3/input.txt")
	raw := strings.Split(string(dat), "\n")

	var total int32
	total = 0
	for _, line := range raw {
		total += parseLine(line)
	}
	fmt.Println("part 1:", total)

	total = 0
	for i := 0; i < len(raw); i += 3 {
		map1 := make(map[rune]bool)
		map2 := make(map[rune]bool)
		map3 := make(map[rune]bool)
		for _, letter := range raw[i] {
			map1[letter] = true
		}
		for _, letter := range raw[i+1] {
			map2[letter] = true
		}
		for _, letter := range raw[i+2] {
			map3[letter] = true
		}
		for k, _ := range map1 {
			if map2[k] && map3[k] {
				total += calcPriority(k)
			}
		}
	}
	fmt.Println("part 2:", total)
}

func parseLine(l string) int32 {
	map1 := make(map[rune]bool)
	map2 := make(map[rune]bool)
	for i, letter := range l {
		if i < len(l)/2 {
			map1[letter] = true
		} else {
			map2[letter] = true
		}
	}
	var duplicate rune
	for _, letter := range l {
		if map1[letter] && map2[letter] {
			duplicate = letter
		}
	}
	return calcPriority(duplicate)
}

func calcPriority(letter rune) int32 {
	if strings.ToUpper(string(letter)) == string(letter) {
		return letter - 'A' + 27
	} else {
		return letter - 'a' + 1
	}
}
