package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day2/input.txt")
	raw := strings.Split(string(dat), "\n")

	total := 0
	for _, line := range raw {
		total += calcRound(line)
	}
	fmt.Println("part 1:", total)

	part2 := 0
	for _, line := range raw {
		part2 += calcRound2(line)
	}
	fmt.Println("part 2:", part2)
}

func calcRound2(l string) int {
	var letterToScore = map[string]int{
		"A": 0, // rock     beats 2, loses to 1
		"B": 1, // paper    beats 0, loses to 2
		"C": 2, // scissors beats 1, loses to 0
		"X": 0,
		"Y": 1,
		"Z": 2,
	}
	op := letterToScore[string(l[0])]
	roundType := string(l[2])
	ans := 0
	switch roundType {
	case "X": // must lose: pick their hand -1
		ans = (op-1+3)%3 + 1
	case "Y": // draw: pick their hand
		ans = 3 + op + 1
	case "Z": // win: pick their hand +1
		ans = 6 + (op+1)%3 + 1
	}
	return ans
}

func calcRound(l string) int {
	var letterToScore = map[string]int{
		"A": 0, // rock     beats 2, loses to 1
		"B": 1, // paper    beats 0, loses to 2
		"C": 2, // scissors beats 1, loses to 0
		"X": 0,
		"Y": 1,
		"Z": 2,
	}
	op := letterToScore[string(l[0])]
	myHand := letterToScore[string(l[2])]
	score := 0
	if myHand == op {
		// draw
		score = 3
	} else if op == (myHand+1)%3 {
		// lose
		score = 0
	} else {
		// win
		score = 6
	}
	return myHand + score + 1
}
