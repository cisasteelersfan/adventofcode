package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2023/day2/input.txt")
	raw := strings.Split(string(dat), "\n")

	sumIds := 0
	for _, line := range raw {
		if isValidGame(line) {
			sumIds += getGameId(line)
		}
	}
	fmt.Println("part 1:", sumIds)

	sumPower := 0
	for _, line := range raw {
		sumPower += getPower(line)
	}
	fmt.Println("part 2:", sumPower)
}

func getPower(s string) int { // Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	s = strings.Split(s, ": ")[1]    // 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	rounds := strings.Split(s, "; ") // ["3 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"]
	m := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, round := range rounds {
		for _, count := range strings.Split(round, ", ") {
			pair := strings.Split(count, " ")
			num, _ := strconv.Atoi(pair[0])
			color := pair[1]
			m[color] = max(m[color], num)
		}
	}
	return m["red"] * m["green"] * m["blue"]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func getGameId(s string) int {
	r := regexp.MustCompile("Game (\\d*)")
	num, _ := strconv.Atoi(r.FindStringSubmatch(s)[1])
	return num
}

func isValidGame(s string) bool { // Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	s = strings.Split(s, ": ")[1]    // 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	rounds := strings.Split(s, "; ") // ["3 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"]
	for _, round := range rounds {
		for _, count := range strings.Split(round, ", ") {
			m := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}
			pair := strings.Split(count, " ")
			num, _ := strconv.Atoi(pair[0])
			color := pair[1]
			m[color] += num
			if m["red"] > 12 || m["green"] > 13 || m["blue"] > 14 {
				return false
			}
		}
	}
	return true
}
