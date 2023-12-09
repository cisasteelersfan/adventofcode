package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2023/day9/input.txt")
	lines := strings.Split(string(dat), "\n")

	total := 0
	for _, line := range lines {
		nums := make([]int, 0)
		for _, num := range strings.Split(line, " ") {
			nums = append(nums, getNum(num))
		}
		total += processLine(nums)
	}
	fmt.Println("part 1:", total)
}

func processLine(nums []int) int {
	if allZeros(nums) {
		return 0
	}
	return nums[0] - processLine(getDiffs(nums))
}

func getDiffs(n []int) []int {
	ans := make([]int, len(n)-1)
	for i := 1; i < len(n); i++ {
		ans[i-1] = n[i] - n[i-1]
	}
	return ans
}

func allZeros(n []int) bool {
	for _, num := range n {
		if num != 0 {
			return false
		}
	}
	return true
}

func getNum(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
