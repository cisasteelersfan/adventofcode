package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	dat, _ := os.ReadFile("2024/day9/input.txt")
	nums := getNums(string(dat))

	// approach: naive. unpack the map, then defrag, then calculate.
	// unpack the map:
	arr := make([]int, 0)
	for i, num := range nums {
		if i%2 == 0 { // it's data
			for j := 0; j < num; j++ {
				arr = append(arr, i/2)
			}
		} else {
			for j := 0; j < num; j++ {
				arr = append(arr, -1)
			}
		}
	}

	// defrag
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] == -1 {
			continue
		}
		left := 0
		for ; arr[left] != -1; left++ {
		}
		if left >= i {
			continue
		}
		arr[left] = arr[i]
		arr[i] = -1
	}

	checksum := 0
	for i, num := range arr {
		if num == -1 {
			break
		}
		checksum += i * num
	}
	fmt.Println("Part 1:", checksum)
}

func getNums(s string) []int {
	nums := make([]int, len(s))
	for i, r := range s {
		nums[i], _ = strconv.Atoi(string(r))
	}
	return nums
}
