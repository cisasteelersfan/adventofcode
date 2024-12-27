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

	part2 := make([][]int, len(lines))
	for i, line := range lines {
		part2[i] = getNums(line)
	}
	maxBananas := 0
	for i := -9; i <= 9; i++ {
		for j := -9; j <= 9; j++ {
			if i+j > 9 || i+j < -9 {
				continue
			}
			for k := -9; k <= 9; k++ {
				if j+k > 9 || j+k < -9 {
					continue
				}
				for l := -9; l <= 9; l++ {
					if k+l > 9 || k+l < -9 {
						continue
					}
					bananas := getBananas(i, j, k, l, part2)
					if bananas > maxBananas {
						maxBananas = bananas
					}
					fmt.Println("max, bananas, ijkl:", maxBananas, bananas, i, j, k, l)
				}
			}
		}
	}
	fmt.Println("part 2:", maxBananas) // 2268 is too low, 5000 too high
}

func getBananas(i, j, k, l int, prices [][]int) int {
	bananas := 0
	for _, sequence := range prices {
		// check for the first i,j,k,l sequence
		curBananas := 0
		for x := 1; x < 2001-3; x++ {
			price := sequence[x]
			diff := price - sequence[x-1]
			if diff == i {
				if j == sequence[x+1]-sequence[x] {
					if k == sequence[x+2]-sequence[x+1] {
						if l == sequence[x+3]-sequence[x+2] {
							curBananas = sequence[x+3]
							break
						}
					}
				}
			}
		}
		bananas += curBananas
	}
	return bananas
}

func getNums(l string) []int {
	num, _ := strconv.Atoi(l)
	nums := make([]int, 2001)
	nums[0] = num % 10
	for i := 1; i < 2001; i++ {
		num = ((num * 64) ^ num) % 16777216
		num = ((num / 32) ^ num) % 16777216
		num = ((num * 2048) ^ num) % 16777216
		nums[i] = num % 10
	}
	return nums
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
