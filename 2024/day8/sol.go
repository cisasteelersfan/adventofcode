package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2024/day8/input.txt")
	m := make(map[Point]Dot)
	antinodes := make(map[Point]bool)
	transmitters := make(map[Dot]map[Point]bool)
	for r, rawRow := range strings.Split(string(dat), "\n") {
		for c, char := range rawRow {
			m[Point{r, c}] = Dot{char}
			if char != '.' {
				_, ok := transmitters[Dot{char}]
				if !ok {
					transmitters[Dot{char}] = make(map[Point]bool)
				}
				transmitters[Dot{char}][Point{r, c}] = true
			}
		}
	}
	// Loop through every board location. For each transmitter, calc distance, and determine if antinode.
	rows, cols := len(strings.Split(string(dat), "\n")), len(strings.Split(string(dat), "\n")[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			isAntinode := false
			// go through all transmitters
			for _, locations := range transmitters {
				for p := range locations {
					rowDiff, colDiff := p.r-r, p.c-c
					if rowDiff == 0 && colDiff == 0 {
						continue
					}
					if locations[Point{p.r + rowDiff, p.c + colDiff}] {
						// it's an antinode!
						fmt.Println("is antinode:", r+rowDiff, c+colDiff)
						antinodes[Point{r, c}] = true
						isAntinode = true
						break
					}
				}
				if isAntinode {
					break
				}
			}
		}
	}
	fmt.Println("part 1:", len(antinodes))
}

type Point struct {
	r, c int
}

type Dot struct {
	r rune
}