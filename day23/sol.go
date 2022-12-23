package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const MaxInt = int(math.MaxUint >> 1)

const (
	north int = 0
	south     = 1
	west      = 2
	east      = 3
)

func main() {
	dat, _ := os.ReadFile("day23/input.txt")
	b := board{make(map[point]*elf), make(map[point]*elf), 0}
	for row, line := range strings.Split(string(dat), "\n") {
		for col, char := range line {
			if char == '#' {
				b.elves[point{row, col}] = &elf{point{row, col}, point{row, col}}
			}
		}
	}
	for i := 0; i < 10; i++ {
		b.simulate()
	}
	fmt.Println("part 1:", b.calculateEmpty())

	b = board{make(map[point]*elf), make(map[point]*elf), 0}
	for row, line := range strings.Split(string(dat), "\n") {
		for col, char := range line {
			if char == '#' {
				b.elves[point{row, col}] = &elf{point{row, col}, point{row, col}}
			}
		}
	}
	rounds := 1
	for ; b.simulate(); rounds++ {
	}
	fmt.Println("part 2:", rounds)
}

type board struct {
	elves          map[point]*elf
	proposals      map[point]*elf
	firstDirection int
}

func (b *board) calculateEmpty() int {
	minR, maxR, minC, maxC := MaxInt, -MaxInt, MaxInt, -MaxInt
	for _, e := range b.elves {
		r, c := e.pos.row, e.pos.col
		if r < minR {
			minR = r
		}
		if r > maxR {
			maxR = r
		}
		if c < minC {
			minC = c
		}
		if c > maxC {
			maxC = c
		}
	}
	return (maxR-minR+1)*(maxC-minC+1) - len(b.elves)
}

func (b *board) simulate() bool {
	b.proposals = make(map[point]*elf)
	for _, e := range b.elves {
		b.updateProposal(e)
	}
	moved := false
	for _, e := range b.elves {
		moved = b.moveToProposal(e) || moved
	}

	b.firstDirection = (b.firstDirection + 1) % 4
	return moved
}

type elf struct {
	pos      point
	proposal point
}

func (b *board) updateProposal(e *elf) {
	e.proposal = e.pos
	r, c := e.pos.row, e.pos.col
	adjacent := []point{{r - 1, c - 1}, {r - 1, c}, {r - 1, c + 1}, {r, c - 1}, {r, c + 1}, {r + 1, c - 1}, {r + 1, c}, {r + 1, c + 1}}
	hasAdjacent := false
	for _, p := range adjacent {
		if _, ok := b.elves[p]; ok {
			hasAdjacent = true
			break
		}
	}
	if !hasAdjacent {
		return
	}
	direction := b.firstDirection
	for i := 0; i < 4; i++ {
		proposed := false
		switch direction {
		case north:
			points := []point{{r - 1, c - 1}, {r - 1, c}, {r - 1, c + 1}}
			foundElf := false
			for _, p := range points {
				if _, ok := b.elves[p]; ok {
					foundElf = true
				}
			}
			if foundElf {
				break
			}
			proposed = true
			proposal := point{r - 1, c}
			if otherElf, clash := b.proposals[proposal]; clash {
				// undo
				otherElf.proposal = otherElf.pos
			} else {
				e.proposal = proposal
				b.proposals[proposal] = e
			}
		case south:
			points := []point{{r + 1, c - 1}, {r + 1, c}, {r + 1, c + 1}}
			foundElf := false
			for _, p := range points {
				if _, ok := b.elves[p]; ok {
					foundElf = true
				}
			}
			if foundElf {
				break
			}
			proposed = true
			proposal := point{e.pos.row + 1, e.pos.col}
			if otherElf, clash := b.proposals[proposal]; clash {
				// undo
				otherElf.proposal = otherElf.pos
			} else {
				e.proposal = proposal
				b.proposals[proposal] = e
			}
		case east:
			points := []point{{r - 1, c + 1}, {r, c + 1}, {r + 1, c + 1}}
			foundElf := false
			for _, p := range points {
				if _, ok := b.elves[p]; ok {
					foundElf = true
				}
			}
			if foundElf {
				break
			}
			proposed = true
			proposal := point{e.pos.row, e.pos.col + 1}
			if otherElf, clash := b.proposals[proposal]; clash {
				// undo
				otherElf.proposal = otherElf.pos
			} else {
				e.proposal = proposal
				b.proposals[proposal] = e
			}
		case west:
			points := []point{{r - 1, c - 1}, {r, c - 1}, {r + 1, c - 1}}
			foundElf := false
			for _, p := range points {
				if _, ok := b.elves[p]; ok {
					foundElf = true
				}
			}
			if foundElf {
				break
			}
			proposed = true
			proposal := point{e.pos.row, e.pos.col - 1}
			if otherElf, clash := b.proposals[proposal]; clash {
				// undo
				otherElf.proposal = otherElf.pos
			} else {
				e.proposal = proposal
				b.proposals[proposal] = e
			}
		}
		direction = (direction + 1) % 4
		if proposed {
			break
		}
	}
}

func (b *board) moveToProposal(e *elf) bool {
	moved := true
	if e.pos == e.proposal {
		moved = false
	}
	oldPos := e.pos
	e.pos = e.proposal
	delete(b.elves, oldPos)
	b.elves[e.pos] = e
	return moved
}

type point struct {
	row, col int
}
