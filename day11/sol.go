package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("day11/input.txt")
	monkeyLines := strings.Split(string(dat), "\n\n")

	monkeys := monkeys{count: len(monkeyLines), monkeys: make(map[int]*monkey)}
	for i, line := range monkeyLines {
		monkeys.monkeys[i] = parseMonkey(line, &monkeys)
	}
	for i := 0; i < 20; i++ {
		monkeys.advance()
	}
	fmt.Println("part 1:", getLevel(monkeys))
}

func getLevel(m monkeys) int {
	inspected := make([]int, len(m.monkeys))
	for i := range m.monkeys {
		inspected[i] = m.monkeys[i].inspectedCount
	}
	sort.Ints(inspected)
	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}

func parseMonkey(l string, p *monkeys) *monkey {
	s := strings.Split(l, "\n")
	m := monkey{}
	m.items = parseItems(s[1])
	m.operation = parseOps(s[2])
	m.test = parseTest(s[3])
	m.trueThrow = parseThrow(s[4])
	m.falseThrow = parseThrow(s[5])
	m.monkeys = p
	return &m
}

func parseItems(s string) []int {
	nums := strings.Split(strings.Split(s, ": ")[1], ", ")
	items := make([]int, len(nums))
	for i, numString := range nums {
		num, _ := strconv.Atoi(numString)
		items[i] = num
	}
	return items
}

func parseOps(s string) func(int) int {
	if "old * old" == strings.Split(s, " = ")[1] {
		return func(i int) int { return i * i }
	}
	op := strings.Split(s, "old ")[1]
	num, _ := strconv.Atoi(strings.Split(op, " ")[1])
	if op[0] == '+' {
		return func(i int) int { return i + num }
	} else if op[0] == '*' {
		return func(i int) int { return i * num }
	}
	panic(fmt.Sprint("uh-oh: op[0]=", op[0]))
}

func parseTest(s string) func(int) bool {
	num, _ := strconv.Atoi(strings.Split(s, "by ")[1])
	return func(i int) bool { return i%num == 0 }
}

func parseThrow(s string) int {
	num, _ := strconv.Atoi(strings.Split(s, "monkey ")[1])
	return num
}

type monkeys struct {
	count   int
	monkeys map[int]*monkey
}

func (m *monkeys) advance() {
	for i := 0; i < len(m.monkeys); i++ {
		m.monkeys[i].processItems()
	}
}

func (m *monkeys) print() {
	for i, m := range m.monkeys {
		fmt.Println(i, "monkey:", *m)
	}
}

type monkey struct {
	monkeys        *monkeys
	items          []int
	inspectedCount int
	operation      func(int) int
	trueThrow      int
	falseThrow     int
	test           func(int) bool
}

func (m *monkey) processItems() {
	if len(m.items) == 0 {
		return
	}
	for _, item := range m.items {
		m.inspectedCount++
		worry := m.operation(item) / 3
		if m.test(worry) {
			m.monkeys.monkeys[m.trueThrow].items = append(m.monkeys.monkeys[m.trueThrow].items, worry)
		} else {
			m.monkeys.monkeys[m.falseThrow].items = append(m.monkeys.monkeys[m.falseThrow].items, worry)
		}
	}
	m.items = make([]int, 0)
}
