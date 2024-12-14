package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

	for i := 0; i < 10000; i++ {
		for i, r := range robots {
			newX := (r.x + r.p.x + width) % width
			newY := (r.y + r.p.y + height) % height
			r.p.x = newX
			r.p.y = newY
			robots[i] = r
		}
		if i > 1000 {
			b, interesting := getBoard(width, height, robots)
			if !interesting {
				continue
			}
			fmt.Printf("iteration: %d\n%s", i+1, b)
			time.Sleep(100 * time.Millisecond)
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

func getBoard(width, height int, robots []Robot) (string, bool) {
	arr := make([][]int, height)
	for row := 0; row < height; row++ {
		arr[row] = make([]int, width)
	}
	for _, r := range robots {
		arr[r.p.y][r.p.x]++
	}
	interesting := false
	s := strings.Builder{}
	for row := 0; row < height; row++ {
		consecutive := 0
		for col := 0; col < width; col++ {
			if arr[row][col] > 0 {
				s.WriteRune('x')
				consecutive++
				if consecutive > 10 {
					interesting = true
				}
			} else {
				s.WriteRune('.')
				consecutive = 0
			}
		}
		s.WriteRune('\n')
	}
	return s.String(), interesting
}

type Point struct {
	x, y int
}

type Robot struct {
	x, y int
	p    Point
}
