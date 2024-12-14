package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	dat, _ := os.ReadFile("2024/day14/" + filename)
	width, height := 101, 103
	if filename == "small.txt" || filename == "tiny.txt" {
		width, height = 11, 7
	}
	raw := strings.Split(string(dat), "\n")
	robots := make([]Robot, len(raw))
	for i, line := range raw {
		x, y, right, down := parse(line)
		robots[i] = Robot{right, down, Point{x, y}}
	}

	for i := 0; i < 100; i++ {
		for i, r := range robots {
			newX := (r.x + r.p.x + width) % width
			newY := (r.y + r.p.y + height) % height
			r.p.x = newX
			r.p.y = newY
			robots[i] = r
		}
	}

	// figure out robots in each quadrant
	quadrants := make(map[int]int)
	midX, midY := width/2, height/2
	for _, r := range robots {
		// check first quadrant
		x, y := r.p.x, r.p.y
		if x < midX && y < midY {
			quadrants[0]++
		} else if x > midX && y < midY {
			quadrants[1]++
		} else if x < midX && y > midY {
			quadrants[2]++
		} else if x > midX && y > midY {
			quadrants[3]++
		}
	}
	ans := 1
	fmt.Println("quadrants:", quadrants)
	for _, count := range quadrants {
		ans *= count
	}
	fmt.Println("part 1:", ans)
}

func parse(s string) (x, y, right, down int) {
	split := strings.Split(s, " ")
	first := strings.Split(split[0], ",")
	second := strings.Split(split[1], ",")
	x = getNum(strings.Split(first[0], "=")[1])
	y = getNum(first[1])
	right = getNum(strings.Split(second[0], "=")[1])
	down = getNum(second[1])
	return
}

func getNum(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

type Point struct {
	x, y int
}

type Robot struct {
	x, y int
	p    Point
}
