package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2023/day4/input.txt")
	raw := strings.Split(string(dat), "\n")

	total := 0
	for _, line := range raw {
		total += processLine(line)
	}
	fmt.Println("part 1:", total)
}

func processLine(s string) int {
	cardNumbers := getCardNumbers(s)
	winningNumbers := getWinningNumbers(s)
	points := 0
	for key := range cardNumbers {
		if _, ok := winningNumbers[key]; ok {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func getCardNumbers(s string) map[int]bool {
	m := make(map[int]bool)
	s = strings.Split(s, ":")[1]
	s = strings.Split(s, "|")[0]
	numStr := strings.Fields(s)
	for _, i := range numStr {
		num, _ := strconv.Atoi(i)
		m[num] = true
	}
	return m
}
func getWinningNumbers(s string) map[int]bool {
	m := make(map[int]bool)
	s = strings.Split(s, "|")[1]
	numStr := strings.Fields(s)
	for _, i := range numStr {
		num, _ := strconv.Atoi(i)
		m[num] = true
	}
	return m
}
