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
	initialRow, initialCol := guard.p.row, guard.p.col
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

	// approach: place 1 obstruction on every blank space. See if he gets in a loop.
	stuckTimes := 0
	for p := range m {
		g := Guard{Point{initialRow, initialCol, false}, 0}
		if p.obstacle || (g.p.row == p.row && g.p.col == p.col) {
			continue
		}
		fmt.Println("trying obstacle:", p)
		// try setting this as obstacle
		mapcopy := copymap(m)
		mapcopy[Point{p.row, p.col, true}] = true
		totalSteps := 0

		visited2 := make(map[Point]int) // value is direction
		for g.isOnMap(mapcopy) {
			totalSteps++
			if facing, ok := visited2[g.p]; ok {
				if facing == g.facing || totalSteps > 10000000 {
					fmt.Println("Stuck! position:", g)
					stuckTimes++
					break
				}
			}
			visited2[g.p] = g.facing
			nextPos := Point{g.p.row, g.p.col, false}
			switch g.facing {
			case 0:
				nextPos.row--
			case 1:
				nextPos.col++
			case 2:
				nextPos.row++
			case 3:
				nextPos.col--
			}
			if mapcopy[Point{nextPos.row, nextPos.col, true}] { // obstacle
				g.facing = (g.facing + 1) % 4 // turn
			} else {
				g.p = nextPos // go forward
			}
		}
	}
	fmt.Println("part 2:", stuckTimes)
}

func copymap(m map[Point]bool) map[Point]bool {
	copy := make(map[Point]bool)
	for k, v := range m {
		copy[Point{k.row, k.col, k.obstacle}] = v
	}
	return copy
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
