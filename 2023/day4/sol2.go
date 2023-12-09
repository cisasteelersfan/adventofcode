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

	cardToWins := make(map[int]int)
	for card, line := range raw {
		cardToWins[card] = processLine(line)
	}
	total := 0
	for card := 0; card < len(raw); card++ {
		total += findCards(cardToWins, card)
	}
	fmt.Println("part 2:", total)
}

func findCards(m map[int]int, card int) int {
	if m[card] == 0 {
		return 1
	}
	tot := 1
	for i := 1; i <= m[card]; i++ {
		tot += findCards(m, card+i)
	}
	return tot
}

func processLine(s string) int {
	cardNumbers := getCardNumbers(s)
	winningNumbers := getWinningNumbers(s)
	inCommon := 0
	for key := range cardNumbers {
		if _, ok := winningNumbers[key]; ok {
			inCommon++
		}
	}
	return inCommon
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
