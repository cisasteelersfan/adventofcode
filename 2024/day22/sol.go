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
