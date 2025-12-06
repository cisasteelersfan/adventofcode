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
