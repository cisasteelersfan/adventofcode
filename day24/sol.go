package main

import (
	"fmt"
	"os"
	"strings"
)

const debug = false

func main() {
	// approach: BFS with a function to check if clobbered
	dat, _ := os.ReadFile("day24/input.txt")

	b := parseBoard(string(dat))
	fmt.Println("part 1:", b.shortestPath())
}

type board struct {
	b          [][]rune
	isOverlap  func(r, c, time int) bool
	rows, cols int
}

func parseBoard(s string) board {
	b := board{}
	funcs := make([]func(r, c, time int) bool, 0)
	split := strings.Split(s, "\n")
	b.rows, b.cols = len(split), len(split[0])
	for r := 1; r < b.rows-1; r++ {
		for c := 1; c < b.cols-1; c++ {
			r := r
			c := c
			run := rune(split[r][c])
			switch run {
			case '>':
				funcs = append(funcs, func(row, col, time int) bool {
					newC := (time%(b.cols-2)+c-1)%(b.cols-2) + 1
					return r == row && newC == col
				})
			case '<':
				funcs = append(funcs, func(row, col, time int) bool {
					newC := (-time%(b.cols-2)+c-1+b.cols-2)%(b.cols-2) + 1
					return r == row && newC == col
				})
			case 'v':
				funcs = append(funcs, func(row, col, time int) bool {
					newR := (time%(b.rows-2)+r-1+b.rows-2)%(b.rows-2) + 1
					return row == newR && c == col
				})
			case '^':
				funcs = append(funcs, func(row, col, time int) bool {
					newR := (-time%(b.rows-2)+r-1+b.rows-2)%(b.rows-2) + 1
					return row == newR && c == col
				})
			}
		}
	}

	b.isOverlap = func(r, c, time int) bool {
		if debug {
			fmt.Println("considering:", r, c, time)
		}
		for _, f := range funcs {
			if f(r, c, time) {
				if debug {
					fmt.Println(r, c, time, "overlaps.")
				}
				return true
			}
		}
		if debug {
			fmt.Println(r, c, time, "NOToverlaps.")
		}
		return false
	}

	return b
}

type node struct {
	r, c, time int
}

func (b *board) isWin(r, c int) bool {
	return r == b.rows-1 && c == b.cols-2
}

func (b *board) getValidNeighbors(r, c, time int) []node {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := make([]node, 0)
	for _, d := range dirs {
		row, col := d[0]+r, d[1]+c
		if row < 1 || col < 1 || row > b.rows-2 || col > b.cols-2 {
			continue
		}
		if !b.isOverlap(row, col, time) {
			n = append(n, node{row, col, 0})
		}
	}
	if r == b.rows-2 && c == b.cols-2 {
		n = append(n, node{b.rows - 1, b.cols - 2, 0})
	}
	if debug {
		fmt.Println("neighbors:", len(n))
	}
	return n
}

func (b *board) shortestPath() int {
	time := 0
	s := make(map[node]bool)
	s[node{0, 1, 0}] = true
	for {
		fmt.Println(time)
		temp := make(map[node]bool)
		for cur := range s {
			r, c := cur.r, cur.c
			if b.isWin(r, c) {
				return time
			}
			for _, neighbor := range b.getValidNeighbors(r, c, time+1) {
				temp[node{neighbor.r, neighbor.c, 0}] = true
			}
			if !b.isOverlap(r, c, time+1) {
				temp[node{r, c, 0}] = true
			}
		}
		time++
		s = temp
	}
}
