package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	readFile := "input.txt"
	dat, _ := os.ReadFile("2025/day8/" + readFile)
	lines := strings.Split(string(dat), "\n")
	nodes := make(map[string]map[string]bool)
	for _, line := range lines {
		nodes[line] = map[string]bool{line: true}
	}
	distanceToNodes := make(map[int]map[string]bool)
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			distanceToNodes[getDistance(lines[i], lines[j])] = map[string]bool{lines[i]: true, lines[j]: true}
		}
	}
	keys := make([]int, 0)
	for k := range distanceToNodes {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	goalConnections := 1000
	if readFile == "small.txt" {
		goalConnections = 10
	}
	for i := 0; i < goalConnections; i++ {
		n := distanceToNodes[keys[i]]
		twoNodes := make([]string, 0)
		for k := range n {
			twoNodes = append(twoNodes, k)
		}
		// add all neighbors of twonodes[1]
		for key := range nodes[twoNodes[1]] {
			nodes[twoNodes[0]][key] = true
			nodes[key][twoNodes[0]] = true
		}
		for key := range nodes[twoNodes[0]] {
			nodes[twoNodes[1]][key] = true
			nodes[key][twoNodes[1]] = true
		}
	}
	// remove duplicates
	newNodes := make(map[string]map[string]bool)
	for _, n := range lines {
		// if n or any of n's neighbors are in newNodes, skip
		alreadyExists := len(newNodes[n]) > 0
		for neighbor := range nodes[n] {
			if len(newNodes[neighbor]) > 0 {
				alreadyExists = true
				break
			}
		}
		if !alreadyExists {
			newNodes[n] = nodes[n]
		}
	}
	sizes := make([]int, 0)
	for _, line := range lines {
		sizes = append(sizes, len(newNodes[line]))
	}
	sort.Ints(sizes)
	ans := sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
	fmt.Println("Part 1:", ans) // 7452 too low
}

func getDistance(i, j string) int {
	spliti := strings.Split(i, ",")
	splitj := strings.Split(j, ",")
	return squared(getNum(spliti[0])-getNum(splitj[0])) +
		squared(getNum(spliti[1])-getNum(splitj[1])) +
		squared(getNum(spliti[2])-getNum(splitj[2]))
}

func squared(i int) int {
	return i * i
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func getNum(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
