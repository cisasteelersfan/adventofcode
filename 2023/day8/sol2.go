package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2023/day8/input.txt")
	split := strings.Split(string(dat), "\n\n")
	instruction := split[0]
	lines := strings.Split(split[1], "\n")

	m := make(map[string][]string)
	r := regexp.MustCompile("(\\w{3})")
	for _, line := range lines {
		matches := r.FindAllString(line, 3)
		m[matches[0]] = []string{matches[1], matches[2]}
	}

	nodes := make(map[string]int)
	for k := range m {
		if k[2] == 'A' {
			nodes[k] = findDistToZ(instruction, k, m)
		}
	}

	fmt.Println("part 2:", leastCommonMultiple(nodes))
}

func leastCommonMultiple(nodes map[string]int) int {
	max := 1
	largest := 0
	for _, num := range nodes {
		max *= num
		if num > largest {
			largest = num
		}
	}
	fmt.Println("max:", max)
	i := 0
	num := largest
	for !allMultiple(nodes, num) {
		i++
		num = i * largest
	}
	return num
}

func allMultiple(nodes map[string]int, i int) bool {
	for _, num := range nodes {
		if i%num != 0 {
			return false
		}
	}
	return true
}

func findDistToZ(instruction, k string, m map[string][]string) int {
	i := 0
	for ; k[2] != 'Z'; i++ {
		direction := 0
		if instruction[i%len(instruction)] == 'R' {
			direction = 1
		}
		k = m[k][direction]
	}
	return i
}
