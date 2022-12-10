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

	ropes := make([]*rope, 10)
	for i := 0; i < len(ropes); i++ {
		ropes[i] = &rope{visited: make(map[point]bool)}
		ropes[i].visited[point{0, 0}] = true
	}
	lr := longRope{ropes}
	for _, line := range raw {
		direction, num := parseLine(line)
		for i := 0; i < num; i++ {
			switch direction {
			case "R":
				lr.ropes[0].head.col++
			case "L":
				lr.ropes[0].head.col--
			case "D":
				lr.ropes[0].head.row--
			case "U":
				lr.ropes[0].head.row++
			}
			lr.moveTails()
			// printBoard(lr)
		}
	}
	fmt.Println("part 2", len(lr.ropes[8].visited)) // 1960, 2205, 2206 are too low
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i, lr.ropes[i].head)
	// }
}

func printBoard(lr longRope) {
	maxRow, maxCol := 0, 0
	minRow, minCol := 0, 0
	for _, r := range lr.ropes {
		if r.head.row < minRow {
			minRow = r.head.row
		}
		if r.head.row > maxRow {
			maxRow = r.head.row
		}
		if r.head.col > maxCol {
			maxCol = r.head.col
		}
		if r.head.col < minCol {
			minCol = r.head.col
		}
	}
	board := make([][]string, maxRow-minRow+1)
	for row := minRow; row < maxRow+1; row++ {
		board[row-minRow] = make([]string, maxCol-minCol+1)
		for col := minCol; col < maxCol+1; col++ {
			wrote := false
			for i := 8; i >= 0; i-- {
				if lr.ropes[i].head.col == col && lr.ropes[i].head.row == row {
					board[row-minRow][col-minCol] = strconv.Itoa(i)
					wrote = true
				}
			}
			if !wrote {
				board[row-minRow][col-minCol] = "."
			}
		}
	}
	fmt.Println("-------")
	for row := maxRow - minRow; row >= 0; row-- {
		fmt.Println(board[row])
	}
}

type point struct {
	row, col int
}

type longRope struct {
	ropes []*rope
}

func (lr *longRope) moveTails() {
	lr.ropes[0].moveTail()
	for i := 1; i < len(lr.ropes); i++ {
		hr := lr.ropes[i-1]
		r := lr.ropes[i]
		r.head.col = hr.tail.col
		r.head.row = hr.tail.row
		r.moveTail()
		if i == 8 {
			// fmt.Println(r.tail.row, r.tail.col)
		}
	}
}

type rope struct {
	head, tail point
	visited    map[point]bool
}

func (r *rope) moveTail() {
	r.maybeMoveDiag(r.tail)
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

func (r *rope) maybeMoveDiag(p point) {
	if p.row == r.head.row+2 && p.col == r.head.col+2 {
		r.tail.row = r.head.row + 1
		r.tail.col = r.head.col + 1
	} else if p.row == r.head.row-2 && p.col == r.head.col-2 {
		r.tail.row = r.head.row - 1
		r.tail.col = r.head.col - 1
	} else if p.row == r.head.row+2 && p.col == r.head.col-2 {
		r.tail.row = r.head.row + 1
		r.tail.col = r.head.col - 1
	} else if p.row == r.head.row-2 && p.col == r.head.col+2 {
		r.tail.row = r.head.row - 1
		r.tail.col = r.head.col + 1
	}
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
