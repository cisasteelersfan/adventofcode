package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2024/day11/input.txt")
	m := make(map[int]int)
	for _, numStr := range strings.Split(string(dat), " ") {
		m[getNum(numStr)]++
	}

	for i := 0; i < 75; i++ {
		m = split(m)
	}
	ans := 0
	for _, count := range m {
		ans += count
	}
	fmt.Println("part 2:", ans)
}

func split(m map[int]int) map[int]int {
	n := make(map[int]int)
	for num, count := range m {
		if num == 0 {
			n[1] += count
		} else if even(num) {
			i, j := splitNum(num)
			n[i] += count
			n[j] += count
		} else {
			n[num*2024] += count
		}
	}
	return n
}

func splitNum(i int) (int, int) {
	numStr := strconv.Itoa(i)
	j, _ := strconv.Atoi(numStr[0 : len(numStr)/2])
	k, _ := strconv.Atoi(numStr[len(numStr)/2:])
	return j, k
}

func even(i int) bool {
	numStr := strconv.Itoa(i)
	return len(numStr)%2 == 0
}

func getNum(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
