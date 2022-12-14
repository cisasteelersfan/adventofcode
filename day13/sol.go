package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day13/input.txt")
	pairs := strings.Split(string(dat), "\n\n")

	part1 := sumPairsInOrder(pairs)
	fmt.Println("part 1:", part1)
}

func sumPairsInOrder(s []string) int {
	total := 0
	for i, line := range s {
		split := strings.Split(line, "\n")
		left, right := split[0], split[1]
		if isInOrder(left, right) {
			total += (i + 1)
		}
	}
	return total
}

func isInOrder(left, right string) bool {
	var l, r any
	json.Unmarshal([]byte(left), &l)
	json.Unmarshal([]byte(right), &r)
	return cmp(l, r) <= 0
}

func cmp(a, b any) int {
	as, aok := a.([]any)
	bs, bok := b.([]any)

	switch {
	case !aok && !bok:
		return int(a.(float64) - b.(float64))
	case !aok:
		as = []any{a}
	case !bok:
		bs = []any{b}
	}

	for i := 0; i < len(as) && i < len(bs); i++ {
		if c := cmp(as[i], bs[i]); c != 0 {
			return c
		}
	}
	return len(as) - len(bs)
}
