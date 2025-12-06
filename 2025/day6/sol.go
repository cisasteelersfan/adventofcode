package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2025/day6/input.txt")
	lines := strings.Split(string(dat), "\n")
	nums := make([][]int, 0)
	ops := make([]string, 0)
	for i, line := range lines[:len(lines)-1] {
		fields := strings.Fields(line)
		nums = append(nums, make([]int, len(fields)))
		for j, s := range fields {
			nums[i][j] = getNum(s)
		}
	}
	for _, f := range strings.Fields(lines[len(lines)-1]) {
		ops = append(ops, f)
	}

	ans := 0
	for i, op := range ops {
		if op == "*" {
			ans += multiply(nums, i)
		} else {
			ans += add(nums, i)
		}
	}
	fmt.Println("Part 1:", ans)

	blankLine := ""
	for range len(lines[0]) {
		blankLine += " "
	}
	lines = append(lines, blankLine)
	lines[len(lines)-1] = lines[len(lines)-2]
	lines[len(lines)-2] = blankLine

	cols := make([]string, 0)
	for col := range len(lines[0]) {
		s := ""
		for row := range len(lines) {
			s += string(lines[row][col])
		}
		cols = append(cols, s)
	}

	ans = 0
	runningProb := 0
	t := "+"
	for _, col := range cols {
		f := strings.Fields(col)
		if len(f) == 0 { // end of problem
			ans += runningProb
			runningProb = 0
			continue
		}
		if f[len(f)-1] == "*" || f[len(f)-1] == "+" {
			t = f[len(f)-1]
			if t == "*" {
				runningProb = 1
			}
			if t == "+" {
				runningProb = 0
			}
		}
		num := getNum(f[0])
		if t == "*" {
			runningProb *= num
		} else {
			runningProb += num
		}
	}
	ans += runningProb
	fmt.Println("Part 2:", ans)
}

func add(nums [][]int, index int) int {
	ans := 0
	for j := range len(nums) {
		ans += nums[j][index]
	}
	return ans
}

func multiply(nums [][]int, index int) int {
	ans := 1
	for j := range len(nums) {
		ans *= nums[j][index]
	}
	return ans
}

func getNum(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
