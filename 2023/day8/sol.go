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
	i := 0
	for l := "AAA"; l != "ZZZ"; i++ {
		direction := 0
		if instruction[i%len(instruction)] == 'R' {
			direction = 1
		}
		l = m[l][direction]
	}
	fmt.Println("part 1:", i)
}
