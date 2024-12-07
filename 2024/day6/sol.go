package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2024/day6/input.txt")
	raw := strings.Split(string(dat), "\n")
	guard := Guard{}
	m := make(map[Point]bool)
	for r, row := range raw {
		for c, col := range row {
			switch col {
			case '.':
				m[Point{r, c, false}] = true
			case '#':
				m[Point{r, c, true}] = true
			case '^':
				m[Point{r, c, false}] = true
				guard = Guard{Point{r, c, false}, 0}
			}
		}
	}
	visited := make(map[Point]bool)

	for guard.isOnMap(m) {
		visited[guard.p] = true
		nextPos := Point{guard.p.row, guard.p.col, false}
		switch guard.facing {
		case 0:
			nextPos.row--
		case 1:
			nextPos.col++
		case 2:
			nextPos.row++
		case 3:
			nextPos.col--
		}
		if m[Point{nextPos.row, nextPos.col, true}] { // obstacle
			guard.facing = (guard.facing + 1) % 4
		} else {
			guard.p = nextPos
		}
	}
	fmt.Println("part 1:", len(visited))
}

type Guard struct {
	p      Point
	facing int // 0 up, 1 right, 2 down, 3 left
}

func (g Guard) isOnMap(m map[Point]bool) bool {
	return m[g.p]
}

type Point struct {
	row, col int
	obstacle bool
}
