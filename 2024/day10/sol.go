package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2024/day10/input.txt")
	// create the graph. apply dfs starting from every 0.
	lookup := make(map[Point]int)
	g := make(map[Point]map[Point]bool) // map from point to point
	rows, cols := len(strings.Split(string(dat), "\n")), len(strings.Split(string(dat), "\n")[0])
	for r, row := range strings.Split(string(dat), "\n") {
		for c, col := range row {
			lookup[Point{r, c}] = getNum(col)
		}
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			num := lookup[Point{r, c}]
			for _, j := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				p := Point{r + j[0], c + j[1]}
				if res, ok := lookup[p]; ok {
					if num+1 == res {
						if _, ok := g[Point{r, c}]; !ok {
							g[Point{r, c}] = make(map[Point]bool)
						}
						g[Point{r, c}][p] = true
					}
				}
			}
		}
	}
	sumScore := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if lookup[Point{r, c}] != 0 {
				continue
			}
			visited := make(map[Point]bool)
			sumScore += walkGraph(g, r, c, visited, lookup)
		}
	}
	fmt.Println("part 2:", sumScore)
}

func walkGraph(g map[Point]map[Point]bool, r, c int, visited map[Point]bool, lookup map[Point]int) int {
	p := Point{r, c}
	// if visited[p] {
	// 	return 0
	// }
	// visited[p] = true
	if lookup[p] == 9 {
		return 1
	}
	num := 0
	for n := range g[p] {
		num += walkGraph(g, n.r, n.c, visited, lookup)
	}
	return num
}

type Point struct {
	r, c int
}

func getNum(r rune) int {
	num, err := strconv.Atoi(string(r))
	if err != nil {
		return -100
	}
	return num
}
