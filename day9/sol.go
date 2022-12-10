package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day9/input.txt")
	raw := strings.Split(string(dat), "\n")
	r := rope{visited: make(map[point]bool)}
	r.visited[point{0, 0}] = true
	for _, line := range raw {
		direction, num := parseLine(line)
		for i := 0; i < num; i++ {
			switch direction {
			case "R":
				r.head.col++
			case "L":
				r.head.col--
			case "D":
				r.head.row--
			case "U":
				r.head.row++
			}
			r.moveTail()
		}
	}
	fmt.Println("part 1:", len(r.visited)) // 6376
}

type point struct {
	row, col int
}

type rope struct {
	head, tail point
	visited    map[point]bool
}

func (r *rope) moveTail() {
	if r.shouldMoveRight(r.tail) {
		r.tail.row = r.head.row
		r.tail.col = r.head.col - 1
	} else if r.shouldMoveLeft(r.tail) {
		r.tail.row = r.head.row
		r.tail.col = r.head.col + 1
	} else if r.shouldMoveDown(r.tail) {
		r.tail.col = r.head.col
		r.tail.row = r.head.row + 1
	} else if r.shouldMoveUp(r.tail) {
		r.tail.col = r.head.col
		r.tail.row = r.head.row - 1
	}
	r.visited[r.tail] = true
}

func (r rope) shouldMoveRight(p point) bool {
	return p.col == r.head.col-2
}
func (r rope) shouldMoveLeft(p point) bool {
	return p.col == r.head.col+2
}
func (r rope) shouldMoveUp(p point) bool {
	return p.row == r.head.row-2
}
func (r rope) shouldMoveDown(p point) bool {
	return p.row == r.head.row+2
}

func parseLine(l string) (string, int) {
	s := strings.Split(l, " ")
	num, _ := strconv.Atoi(s[1])
	return s[0], num
}
