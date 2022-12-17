package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day15/input.txt")
	lines := strings.Split(string(dat), "\n")

	pairs := make([]pair, len(lines))
	for i, line := range lines {
		pairs[i] = parseLine(line)
	}
	// fmt.Println("part 1:", findNoBeacons(pairs, 10))
	fmt.Println("part 1:", findNoBeacons(pairs, 2000000))
	// fmt.Println("part 2:", findFrequency(pairs, 20))
	fmt.Println("part 2:", findFrequency(pairs, 4000000))
}

func findFrequency(pairs []pair, max int) int {
	// look one outside the perimeter of each square.
	for _, p := range pairs {
		for _, posPoint := range p.getPerimeter(max) {
			couldBePos := true
			for _, otherPair := range pairs {
				if getManhattan(posPoint, otherPair.sensor) <= otherPair.getManhattan() {
					couldBePos = false
					break
				}
			}
			if couldBePos {
				// fmt.Println("x", posPoint.x, "y", posPoint.y)
				return posPoint.y + posPoint.x*4000000
			}
		}
	}
	panic("didn't find a possible beacon location")
}

func (p pair) getPerimeter(max int) []point {
	perimeter := make([]point, 0)
	m := p.getManhattan()
	for posX := -2 * m; posX <= 2*m; posX++ {
		var x, y int
		if posX == -2*m || posX == 2*m {
			x, y = p.sensor.x+posX, p.sensor.y
			if x < 0 || y < 0 || x > max || y > max {
				continue
			}
			perimeter = append(perimeter, point{x, y})
		} else {
			posY := m - posX + 1
			x, y = p.sensor.x+posX, posY+p.sensor.y
			if x < 0 || y < 0 || x > max || y > max {
				continue
			}
			perimeter = append(perimeter, point{x, y})
			y = p.sensor.y - posY
			if y < 0 || y > max {
				continue
			}
			perimeter = append(perimeter, point{x, y})
		}
	}
	// fmt.Println("perimeter", perimeter)
	return perimeter
}

func findNoBeacons(pairs []pair, row int) int {
	noBeacons := 0
	minX, maxX := getMinMax(pairs, row)
	for x := minX; x <= maxX; x++ {
		withinDist := false
		for _, p := range pairs {
			if p.beacon.y == row && p.beacon.x == x {
				continue
			}
			if p.withinManhattan(x, row) {
				withinDist = true
			}
		}
		if withinDist {
			noBeacons++
		}
	}
	return noBeacons
}

func (p pair) withinManhattan(x, y int) bool {
	dist := getManhattan(p.sensor, point{x, y})
	return dist <= p.getManhattan()
}

func getMinMax(pairs []pair, row int) (int, int) {
	maxManhattan := 0
	for _, p := range pairs {
		if p.getManhattan() > maxManhattan {
			maxManhattan = p.getManhattan()
		}
	}
	minX, maxX := 0, 0
	for _, p := range pairs {
		if p.beacon.x < minX {
			minX = p.beacon.x
		} else if p.beacon.x > maxX {
			maxX = p.beacon.x
		} else if p.sensor.x < minX {
			minX = p.sensor.x
		} else if p.sensor.x > maxX {
			maxX = p.sensor.x
		}
	}
	return minX - maxManhattan, maxX + maxManhattan
}

func parseLine(s string) pair {
	rx := regexp.MustCompile("x=(\\-?\\d+)")
	ry := regexp.MustCompile("y=(\\-?\\d+)")
	p := pair{}
	p.sensor.x = getNum(rx.FindAllStringSubmatch(s, 2)[0][1])
	p.beacon.x = getNum(rx.FindAllStringSubmatch(s, 2)[1][1])
	p.sensor.y = getNum(ry.FindAllStringSubmatch(s, 2)[0][1])
	p.beacon.y = getNum(ry.FindAllStringSubmatch(s, 2)[1][1])
	return p
}

func getNum(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

type pair struct {
	sensor point
	beacon point
}

func (p pair) getManhattan() int {
	return getManhattan(p.sensor, p.beacon)
}

func getManhattan(s, d point) int {
	return int(math.Abs(float64(s.x)-float64(d.x)) +
		math.Abs(float64(s.y)-float64(d.y)))
}

type point struct {
	x, y int
}
