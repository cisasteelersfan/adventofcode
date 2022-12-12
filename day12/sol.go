package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day12/input.txt")
	g := parseGraph(string(dat))
	start, end := getStart(g), getEnd(g)
	fmt.Println("part 1:", findShortestPaths(g, start, end))
}

func findShortestPaths(g graph, start, end point) int {
	visited := make(map[point]bool)
	visited[start] = true
	q := make([]point, 1)
	q[0] = start
	g.nodes[start].dist = 0
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		if v == end {
			return g.nodes[end].dist
		}
		for n := range g.nodes[v].adjNodes {
			if visited[n] {
				continue
			}
			visited[n] = true
			g.nodes[n].dist = g.nodes[v].dist + 1
			q = append(q, n)
		}
	}
	panic("Didn't encounter end")
}

func getStart(g graph) point {
	for p, n := range g.nodes {
		if n.height == -1 {
			return p
		}
	}
	panic("didn't find height -1")
}

func getEnd(g graph) point {
	for p, n := range g.nodes {
		if n.height == 26 {
			return p
		}
	}
	panic("didn't find height 26")
}

func parseGraph(s string) graph {
	l := strings.Split(s, "\n")
	rows, cols := len(l), len(l[0])
	g := graph{nodes: make(map[point]*node)}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			p := point{r, c}
			height := parseHeight(l[r][c])
			g.addNode(p, height)
		}
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			source := point{r, c}
			dest := point{r + 1, c}
			if g.nodes[dest] != nil && g.nodes[dest].height <= g.nodes[source].height+1 {
				g.addEdge(source, dest)
			}
			dest = point{r - 1, c}
			if g.nodes[dest] != nil && g.nodes[dest].height <= g.nodes[source].height+1 {
				g.addEdge(source, dest)
			}
			dest = point{r, c + 1}
			if g.nodes[dest] != nil && g.nodes[dest].height <= g.nodes[source].height+1 {
				g.addEdge(source, dest)
			}
			dest = point{r, c - 1}
			if g.nodes[dest] != nil && g.nodes[dest].height <= g.nodes[source].height+1 {
				g.addEdge(source, dest)
			}
		}
	}
	return g
}

func parseHeight(r byte) int {
	if r == 'S' {
		return -1
	}
	if r == 'E' {
		return 26
	}
	return int(r - 'a')
}

type graph struct {
	nodes map[point]*node
}

func (g *graph) addNode(key point, height int) {
	if _, found := g.nodes[key]; found {
		return
	}
	g.nodes[key] = &node{p: key, height: height, adjNodes: make(map[point]struct{}), dist: -1}
}

func (g *graph) addEdge(src, dst point) {
	g.nodes[src].adjNodes[dst] = struct{}{}
}

type point struct {
	row, col int
}

type node struct {
	p        point
	height   int
	adjNodes map[point]struct{}
	dist     int
}
