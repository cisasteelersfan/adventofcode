package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2024/day7/input.txt")
	rawLines := strings.Split(string(dat), "\n")

	part1 := 0
	for _, line := range rawLines {
		s := strings.Split(line, ": ")
		test, equation := getNum(s[0]), getNums(s[1])
		if isPossible(test, equation) {
			part1 += test
		}
	}
	fmt.Println("part 1:", part1)

	part2 := 0
	for _, line := range rawLines {
		s := strings.Split(line, ": ")
		test, equation := getNum(s[0]), getNums(s[1])
		if isPossible2(test, equation) {
			part2 += test
		}
	}
	fmt.Println("part 2:", part2)
}

func isPossible2(test int, equation []int) bool {
	if len(equation) == 1 {
		return test == equation[0]
	}
	cur := equation[0]
	next := equation[1]
	add := append([]int{cur + next}, equation[2:]...)
	mul := append([]int{cur * next}, equation[2:]...)
	cat := append([]int{concat(cur, next)}, equation[2:]...)
	return isPossible2(test, add) || isPossible2(test, mul) || isPossible2(test, cat)
}

func concat(i, j int) int {
	is, js := strconv.Itoa(i), strconv.Itoa(j)
	return getNum(is + js)
}

func isPossible(test int, equation []int) bool {
	if len(equation) == 1 {
		return test == equation[0]
	}
	cur := equation[0]
	next := equation[1]
	add := append([]int{cur + next}, equation[2:]...)
	mul := append([]int{cur * next}, equation[2:]...)
	return isPossible(test, add) || isPossible(test, mul)
}

func getNum(s string) int {
	numstr, _ := strconv.Atoi(s)
	return numstr
}

func getNums(s string) []int {
	nums := make([]int, 0)
	for _, numstr := range strings.Split(s, " ") {
		nums = append(nums, getNum(numstr))
	}
	return nums
}
