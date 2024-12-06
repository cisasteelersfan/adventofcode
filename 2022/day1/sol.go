package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	maxCount := 0
	dat, _ := os.ReadFile("day1/input.txt")
	buckets := strings.Split(string(dat), "\n\n")
	calories := []int{}
	for _, bucket := range buckets {
		c := strings.Split(bucket, "\n")
		sumC := sum(c)
		if sumC > maxCount {
			maxCount = sumC
		}
		calories = append(calories, sumC)
	}
	fmt.Println("part 1:", maxCount)

	sort.Ints(calories)
	maxThreeCount := calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3]
	fmt.Println("port 2:", maxThreeCount)
}

func sum(calories []string) int {
	total := 0
	for _, line := range calories {
		n, _ := strconv.Atoi(line)
		total += n
	}
	return total
}
