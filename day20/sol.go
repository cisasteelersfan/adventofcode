package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day20/input.txt")

	d := dll{make([]*node, 0)}
	for _, numStr := range strings.Split(string(dat), "\n") {
		num, _ := strconv.Atoi(numStr)
		d.original = append(d.original, &node{num: num})
	}
	for i := 1; i < len(d.original); i++ {
		d.original[i-1].next = d.original[i]
		d.original[i].prev = d.original[i-1]
	}
	d.original[0].prev = d.original[len(d.original)-1]
	d.original[len(d.original)-1].next = d.original[0]

	for _, n := range d.original {
		d.advance(n)
	}
	fmt.Println("part 1:", d.sum())

	for i := 1; i < len(d.original); i++ {
		d.original[i-1].next = d.original[i]
		d.original[i].prev = d.original[i-1]
	}
	d.original[0].prev = d.original[len(d.original)-1]
	d.original[len(d.original)-1].next = d.original[0]
	for _, n := range d.original {
		n.num *= 811589153
	}
	for i := 0; i < 10; i++ {
		for _, n := range d.original {
			d.advance(n)
		}
	}
	fmt.Println("part 2:", d.sum())
}

func (d *dll) sum() int {
	zero := &node{}
	for _, n := range d.original {
		if n.num == 0 {
			zero = n
			break
		}
	}
	oneT, twoT, threeT := 0, 0, 0
	next := zero
	for i := 1; i <= 3000; i++ {
		next = next.next
		if i == 1000 {
			oneT = next.num
		} else if i == 2000 {
			twoT = next.num
		} else if i == 3000 {
			threeT = next.num
		}
	}
	return oneT + twoT + threeT
}

func (d *dll) print() {
	start := d.original[0]
	for range d.original {
		fmt.Println(start.num)
		start = start.next
	}
	fmt.Println()
}

func (d *dll) advance(n *node) {
	num := n.num
	if num > 0 {
		num = num % (len(d.original) - 1)
		n.prev.next = n.next
		n.next.prev = n.prev
		last := n.prev
		for i := 0; i < num; i++ {
			last = last.next
		}
		n.next = last.next
		last.next.prev = n
		last.next = n
		n.prev = last
	} else if num < 0 {
		num = num % (len(d.original) - 1)
		n.prev.next = n.next
		n.next.prev = n.prev
		last := n.next
		for i := 0; i < -num; i++ {
			last = last.prev
		}
		n.prev = last.prev
		n.prev.next = n
		n.next = last
		last.prev = n
	}
}

type dll struct {
	original []*node
}

type node struct {
	num        int
	prev, next *node
}
