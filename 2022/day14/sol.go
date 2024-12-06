package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day14/input.txt")
	m := make(map[point]bool)
	largestY := 0
	for _, line := range strings.Split(string(dat), "\n") {
		points := strings.Split(line, " -> ")
		for i := 1; i < len(points); i++ {
			source := parsePoint(points[i-1])
			dest := parsePoint(points[i])
			for _, point := range getPoints(source, dest) {
				m[point] = true
				if point.y > largestY {
					largestY = point.y
				}
			}
		}
	}
	copy := make(map[point]bool)
	for k, v := range m {
		copy[k] = v
	}
	drops := 0
	for i := 0; simulateDrop(500, -1000, m); i++ {
		drops++
	}
	fmt.Println("part 1:", drops)

	drops = 0
	for i := 0; !copy[point{500, 0}]; i++ {
		drops++
		simulateDropWithFloor(500, 0, copy, largestY)
	}
	fmt.Println("part 2:", drops)
}

func simulateDropWithFloor(x, y int, m map[point]bool, largestY int) {
	if y == largestY+1 {
		m[point{x, y}] = true
		return
	}
	if m[point{x, y + 1}] { // if below is occupied
		if !m[point{x - 1, y + 1}] {
			simulateDropWithFloor(x-1, y+1, m, largestY)
			return
		}
		if !m[point{x + 1, y + 1}] {
			simulateDropWithFloor(x+1, y+1, m, largestY)
			return
		}
		m[point{x, y}] = true
		return
	}
	simulateDropWithFloor(x, y+1, m, largestY)
	return
}

// returns true if sand was placed
func simulateDrop(x, y int, m map[point]bool) bool {
	if y > 10000 {
		return false
	}
	if m[point{x, y + 1}] { // if below is occupied
		if !m[point{x - 1, y + 1}] {
			return simulateDrop(x-1, y+1, m)
		}
		if !m[point{x + 1, y + 1}] {
			return simulateDrop(x+1, y+1, m)
		}
		m[point{x, y}] = true
		return true
	}
	return simulateDrop(x, y+1, m)
}

func getPoints(s, d point) []point {
	points := make([]point, 0)
	if s.x == d.x {
		if s.y > d.y {
			s, d = d, s
		}
		for y := s.y; y <= d.y; y++ {
			points = append(points, point{s.x, y})
		}
	} else if s.y == d.y {
		if s.x > d.x {
			s, d = d, s
		}
		for x := s.x; x <= d.x; x++ {
			points = append(points, point{x, s.y})
		}
	}
	return points
}

func parsePoint(s string) point {
	split := strings.Split(s, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return point{x, y}
}

type point struct {
	x, y int
}
