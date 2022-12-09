package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day7/input.txt")
	raw := strings.Split(string(dat), "\n")

	g := createGraph(raw)

	sum := getSumDirsUnder(g, 100_000)
	fmt.Println("part 1:", sum)
}

type node struct {
	isDir    bool
	name     string
	size     int
	children []*node
	parent   *node
}

func getSumDirsUnder(g *node, under int) int {
	if g == nil {
		return 0
	}
	ans := 0
	if g.isDir {
		if g.size < under {
			ans = g.size
		}
	}
	for _, n := range g.children {
		ans += getSumDirsUnder(n, under)
	}
	return ans
}

func (n *node) addChild(child *node) {
	if n.children == nil {
		n.children = make([]*node, 0)
	}
	child.parent = n
	n.children = append(n.children, child)
	n.size += child.size
	for p := n.parent; p != nil; p = p.parent {
		p.size += child.size
	}
}

func createGraph(inst []string) *node {
	root := node{isDir: true, name: "/"}
	curNode := &root
	for i := 1; i < len(inst); i++ {
		if isCdUp(inst[i]) {
			curNode = curNode.parent
		} else if isCd(inst[i]) {
			n := node{isDir: true, name: getName(inst[i])}
			curNode.addChild(&n)
			curNode = &n
		} else if isFile(inst[i]) {
			n := node{isDir: false, name: getName(inst[i]), size: getSize(inst[i])}
			curNode.addChild(&n)
		}
	}
	return &root
}

func isCdUp(s string) bool {
	return s == "$ cd .."
}

func getSize(s string) int {
	raw := strings.Split(s, " ")[0]
	num, _ := strconv.Atoi(raw)
	return num
}

func isFile(s string) bool {
	match, _ := regexp.MatchString("^\\d+", s)
	return match
}

func getName(s string) string {
	r := regexp.MustCompile("(?:$ )?\\w+ (.*)")
	found := r.FindStringSubmatch(s)[1]
	return found
}

func isCd(s string) bool {
	return s[0:4] == "$ cd"
}
