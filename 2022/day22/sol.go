package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day22/input.txt")
	b := &board{b: make(map[point]rune)}

	for row, line := range strings.Split(strings.Split(string(dat), "\n\n")[0], "\n") {
		if row > b.maxRows {
			b.maxRows = row
		}
		for col, char := range line {
			if col > b.maxCols {
				b.maxCols = col
			}
			switch char {
			case '#':
				b.b[point{row, col}] = '#'
			case '.':
				b.b[point{row, col}] = '.'
			}
		}
	}
	for col := 0; col < b.maxCols; col++ {
		if b.b[point{0, col}] == '.' {
			b.pos.p = point{0, col}
			break
		}
	}

	nums, turns := tokenize(strings.Split(string(dat), "\n\n")[1])
	for j := 0; j < nums[0]; j++ {
		b.advance()
	}
	for i := 0; i < len(turns); i++ {
		b.turn(turns[i])
		for j := 0; j < nums[i+1]; j++ {
			b.advance()
		}
	}
	fmt.Println("part 1:", 1000*(b.pos.p.row+1)+4*(b.pos.p.col+1)+b.pos.facing)
}

func (b *board) advance() {
	nextCandidate := point{}
	switch b.pos.facing {
	case right:
		_, found := b.b[point{b.pos.p.row, b.pos.p.col + 1}]
		if found {
			nextCandidate = point{b.pos.p.row, b.pos.p.col + 1}
		} else { // wrap around
			for col := 0; col <= b.pos.p.col; col++ {
				_, found := b.b[point{b.pos.p.row, col}]
				if found {
					nextCandidate = point{b.pos.p.row, col}
					break
				}
			}
		}
	case left:
		_, found := b.b[point{b.pos.p.row, b.pos.p.col - 1}]
		if found {
			nextCandidate = point{b.pos.p.row, b.pos.p.col - 1}
		} else { // wrap around
			for col := b.maxCols; col >= 0; col-- {
				_, found := b.b[point{b.pos.p.row, col}]
				if found {
					nextCandidate = point{b.pos.p.row, col}
					break
				}
			}
		}
	case up:
		_, found := b.b[point{b.pos.p.row - 1, b.pos.p.col}]
		if found {
			nextCandidate = point{b.pos.p.row - 1, b.pos.p.col}
		} else { // wrap down
			for row := b.maxRows; row >= 0; row-- {
				_, found := b.b[point{row, b.pos.p.col}]
				if found {
					nextCandidate = point{row, b.pos.p.col}
					break
				}
			}
		}
	case down:
		_, found := b.b[point{b.pos.p.row + 1, b.pos.p.col}]
		if found {
			nextCandidate = point{b.pos.p.row + 1, b.pos.p.col}
		} else { // wrap up
			for row := 0; row <= b.maxRows; row++ {
				_, found := b.b[point{row, b.pos.p.col}]
				if found {
					nextCandidate = point{row, b.pos.p.col}
					break
				}
			}
		}
	}
	if b.b[nextCandidate] == '.' {
		b.pos.p = nextCandidate
	}
}

func (b *board) turn(r rune) {
	if r == 'R' {
		b.pos.facing = (b.pos.facing + 1) % 4
	} else {
		b.pos.facing = (b.pos.facing + 3) % 4
	}
}

func tokenize(s string) ([]int, []rune) {
	replaced := strings.ReplaceAll(s, "L", "R")
	nums := strings.Split(replaced, "R")
	turns := make([]rune, 0)
	for _, r := range s {
		if r == 'R' || r == 'L' {
			turns = append(turns, r)
		}
	}
	numsConv := make([]int, len(nums))
	for i, num := range nums {
		n, _ := strconv.Atoi(num)
		numsConv[i] = n
	}
	return numsConv, turns
}

const (
	right int = 0
	down      = 1
	left      = 2
	up        = 3
)

type position struct {
	p      point
	facing int
}

type point struct {
	row, col int
}

type board struct {
	pos              position
	b                map[point]rune
	maxRows, maxCols int
}
