package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day5/input.txt")
	before, after, _ := strings.Cut(string(dat), "\n\n")

	stacks := parseListofStacks(before)
	moves := strings.Split(after, "\n")
	for _, line := range moves {
		performMove(stacks, line, false)
	}
	fmt.Println("part 1:", getTop(stacks))

	stacks = parseListofStacks(before)
	moves = strings.Split(after, "\n")
	for _, line := range moves {
		performMove(stacks, line, true)
	}
	fmt.Println("part 2:", getTop(stacks))
}

type stack struct {
	internal []byte
}

func (s *stack) push(r byte) {
	if len(s.internal) == 0 {
		s.internal = make([]byte, 0)
	}
	s.internal = append(s.internal, r)
}

func (s *stack) pop() byte {
	if len(s.internal) > 0 {
		popped := s.internal[len(s.internal)-1]
		s.internal = s.internal[0 : len(s.internal)-1]
		return popped
	}
	panic("tried to pop empty stack")
}

func parseListofStacks(s string) []stack {
	lines := strings.Split(s, "\n")
	numStacks := (len(lines[0]) + 1) / 4
	stacks := make([]stack, numStacks)

	for _, l := range lines[0 : len(lines)-1] {
		pushOnStacks(stacks, l)
	}
	stacks = reverseStacks(stacks)
	return stacks
}

func pushOnStacks(stacks []stack, l string) {
	for i, idx := 0, 1; idx < len(l); i, idx = i+1, idx+4 {
		if string(l[idx]) != " " {
			stacks[i].push(l[idx])
		}
	}
}

func reverseStacks(s []stack) []stack {
	newStacks := make([]stack, len(s))
	for i, oldStack := range s {
		for len(oldStack.internal) > 0 {
			newStacks[i].push(oldStack.pop())
		}
	}
	return newStacks
}

func performMove(stacks []stack, line string, inOrder bool) {
	split := strings.Split(line, " ")
	count, from, to := toInt(split[1]), toInt(split[3])-1, toInt(split[5])-1
	if inOrder {
		tempStack := stack{}
		for i := 0; i < count; i++ {
			tempStack.push(stacks[from].pop())
		}
		for i := 0; i < count; i++ {
			stacks[to].push(tempStack.pop())
		}
	} else {
		for i := 0; i < count; i++ {
			stacks[to].push(stacks[from].pop())
		}
	}
}

func getTop(stacks []stack) string {
	ans := make([]string, len(stacks))
	for i, s := range stacks {
		ans[i] = string(s.pop())
	}
	return strings.Join(ans, "")
}

func toInt(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}
