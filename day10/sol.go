package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day10/input.txt")
	lines := strings.Split(string(dat), "\n")
	clockToRegister := []int{1}
	prevRegVal := 1
	for _, line := range lines {
		operation, value := parseLine(line)
		if operation == "noop" {
			clockToRegister = append(clockToRegister, prevRegVal)
		} else {
			newVal := prevRegVal + value
			clockToRegister = append(clockToRegister, prevRegVal)
			clockToRegister = append(clockToRegister, newVal)
			prevRegVal = newVal
		}
	}
	fmt.Println("part 1:", getSumStrengths(clockToRegister))
}

func getSumStrengths(clockToRegister []int) int {
	sum := 0
	for i := 19; i <= 219; i += 40 {
		sum += (i + 1) * clockToRegister[i]
	}
	return sum
}

func parseLine(l string) (string, int) {
	s := strings.Split(l, " ")
	if len(s) == 1 {
		return l, 0
	}
	num, _ := strconv.Atoi(s[1])
	return s[0], num
}
