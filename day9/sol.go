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
	r := rope{visited: make(map[point]int)}
	r.visited[point{0, 0}]++
	for _, line := range raw {
		direction, num := parseLine(line)
		for i := 0; i < num; i++ {
			switch direction {
			case "R":
				r.right()
			case "L":
				r.left()
			case "D":
				r.down()
			case "U":
				r.up()
			}
		}
	}
	fmt.Println("part 1:", len(r.visited))
}

type point struct {
	row, col int
}

type rope struct {
	head, tail point
	visited    map[point]int
}

func (r *rope) up() {
	r.head.row++
	if r.head.row == r.tail.row+2 {
		r.tail.row = r.head.row - 1
		r.tail.col = r.head.col
		r.visited[r.tail]++
	}
}

func (r *rope) down() {
	r.head.row--
	if r.head.row == r.tail.row-2 {
		r.tail.row = r.head.row + 1
		r.tail.col = r.head.col
		r.visited[r.tail]++
	}
}

func (r *rope) left() {
	r.head.col--
	if r.head.col == r.tail.col-2 {
		r.tail.col = r.head.col + 1
		r.tail.row = r.head.row
		r.visited[r.tail]++
	}
}

func (r *rope) right() {
	r.head.col++
	if r.head.col == r.tail.col+2 {
		r.tail.col = r.head.col - 1
		r.tail.row = r.head.row
		r.visited[r.tail]++
	}
}

func parseLine(l string) (string, int) {
	s := strings.Split(l, " ")
	num, _ := strconv.Atoi(s[1])
	return s[0], num
}
