package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2024/day22/input.txt")
	lines := strings.Split(string(dat), "\n")

	sum := 0
	for _, line := range lines {
		sum += get2000th(line)
	}
	fmt.Println("part 1:", sum)

	// want a list of maps from sequence: price
	part2 := make([]map[Sequence]int, len(lines))
	// map from sequence to total
	tot := make(map[Sequence]int)
	for i, line := range lines {
		part2[i] = getNums(line, tot)
	}
	maxBananas := 0
	for _, val := range tot {
		if val > maxBananas {
			maxBananas = val
		}
	}
	fmt.Println("part 2:", maxBananas) // 2268 is too low, 5000 too high
}

type Sequence struct {
	i, j, k, l int
}

func getNums(l string, tot map[Sequence]int) map[Sequence]int {
	num, _ := strconv.Atoi(l)
	sequenceToPrice := make(map[Sequence]int)
	nums := make([]int, 2001)
	nums[0] = num % 10
	for i := 1; i < 2001; i++ {
		num = ((num * 64) ^ num) % 16777216
		num = ((num / 32) ^ num) % 16777216
		num = ((num * 2048) ^ num) % 16777216
		nums[i] = num % 10
		if i < 4 {
			continue
		}
		s := Sequence{nums[i-4] - nums[i-3], nums[i-3] - nums[i-2], nums[i-2] - nums[i-1], nums[i-1] - nums[i]}
		if _, ok := sequenceToPrice[s]; !ok {
			sequenceToPrice[s] = nums[i]
			tot[s] += nums[i]
		}
	}
	return sequenceToPrice
}

func get2000th(l string) int {
	num, _ := strconv.Atoi(l)
	for i := 0; i < 2000; i++ {
		num = ((num * 64) ^ num) % 16777216
		num = ((num / 32) ^ num) % 16777216
		num = ((num * 2048) ^ num) % 16777216
	}
	return num
}
