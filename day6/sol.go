package main

import (
	"fmt"
	"os"
)

func main() {
	dat, _ := os.ReadFile("day6/input.txt")
	raw := string(dat)

	part1 := findFirstDistinct(raw, 4)

	fmt.Println("part 1:", part1)

	fmt.Println("part 2:", findFirstDistinct(raw, 14))
}

type set struct {
	m map[byte]int
}

func (s *set) add(r byte) {
	s.m[r]++
}

func (s *set) remove(r byte) {
	s.m[r]--
	if s.m[r] == 0 {
		delete(s.m, r)
	}
}

func (s *set) allDistinct() bool {
	for _, v := range s.m {
		if v > 1 {
			return false
		}
	}
	return true
}

func findFirstDistinct(line string, num int) int {
	s := set{make(map[byte]int)}
	for i := 0; i < num; i++ {
		s.add(line[i])
	}
	if s.allDistinct() {
		return num
	}
	for i := num; i < len(line); i++ {
		s.add(line[i])
		s.remove(line[i-num])
		if s.allDistinct() {
			return i + 1
		}
	}
	panic("didn't find distinct.")
}
