package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2025/day2/input.txt")
	blocks := strings.Split(string(dat), ",")
	ranges := make([]Range, 0)
	for _, block := range blocks {
		s := strings.Split(block, "-")
		start, _ := strconv.Atoi(s[0])
		end, _ := strconv.Atoi(s[1])
		r := Range{Start: start, End: end}
		ranges = append(ranges, r)
	}

	ans := 0
	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			s := strconv.Itoa(i)
			if len(s)%2 == 0 {
				first := s[0 : len(s)/2]
				sec := s[len(s)/2:]
				if first == sec {
					ans += i
				}
			}
		}
	}
	fmt.Println("Part 1:", ans)

	ans = 0
	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			if isInvalid(i) {
				ans += i
			}
		}
	}
	fmt.Println("Part 2:", ans)
}

func isInvalid(i int) bool {
	s := strconv.Itoa(i)
	for size := 1; size < len(s); size++ {
		sizeInvalid := true
		if len(s)%size != 0 {
			continue
		}
		first := s[0:size]
		for j := 1; j < len(s)/size; j++ {
			if first != s[j*size:j*size+size] {
				// this SIZE is NOT invalid, but other sizes might be.
				sizeInvalid = false
			}
		}
		if sizeInvalid {
			return true
		}
	}
	return false
}

type Range struct {
	Start int
	End   int
}
