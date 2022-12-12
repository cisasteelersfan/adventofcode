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

	m := monkeys{count: len(monkeyLines), monkeys: make(map[int]*monkey)}
	for i, line := range monkeyLines {
		m.monkeys[i] = parseMonkey(line, &m, false)
	}
	for i := 0; i < 20; i++ {
		m.advance()
	}
	fmt.Println("part 1:", getLevel(m))

	m = monkeys{count: len(monkeyLines), monkeys: make(map[int]*monkey)}
	for i, line := range monkeyLines {
		m.monkeys[i] = parseMonkey(line, &m, true)
	}
	m.worrymod = m.calcWorrymod()
	for i := 0; i < 10000; i++ {
		m.advance()
	}
	fmt.Println("part 2:", getLevel(m))
}

func getLevel(m monkeys) int {
	inspected := make([]int, len(m.monkeys))
	for i := range m.monkeys {
		inspected[i] = m.monkeys[i].inspectedCount
	}
	sort.Ints(inspected)
	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}

func parseMonkey(l string, p *monkeys, part2 bool) *monkey {
	s := strings.Split(l, "\n")
	m := monkey{}
	m.items = parseItems(s[1])
	m.operation = parseOps(s[2])
	m.testNum = parseTest(s[3])
	m.trueThrow = parseThrow(s[4])
	m.falseThrow = parseThrow(s[5])
	m.monkeys = p
	m.part2 = part2
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

func parseTest(s string) int {
	num, _ := strconv.Atoi(strings.Split(s, "by ")[1])
	return num
}

func parseThrow(s string) int {
	num, _ := strconv.Atoi(strings.Split(s, "monkey ")[1])
	return num
}

type monkeys struct {
	count    int
	monkeys  map[int]*monkey
	worrymod int
}

func (m *monkeys) calcWorrymod() int {
	mod := 1
	for _, monkey := range m.monkeys {
		mod *= monkey.testNum
	}
	return mod
}

func (m *monkeys) advance() {
	for i := 0; i < len(m.monkeys); i++ {
		m.monkeys[i].processItems(m.worrymod)
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
	testNum        int
	part2          bool
}

func (m *monkey) processItems(worrymod int) {
	if len(m.items) == 0 {
		return
	}
	for _, item := range m.items {
		m.inspectedCount++
		var worry int
		if m.part2 {
			worry = m.operation(item) % worrymod
		} else {
			worry = m.operation(item) / 3
		}
		if worry%m.testNum == 0 {
			m.monkeys.monkeys[m.trueThrow].items = append(m.monkeys.monkeys[m.trueThrow].items, worry)
		} else {
			m.monkeys.monkeys[m.falseThrow].items = append(m.monkeys.monkeys[m.falseThrow].items, worry)
		}
	}
	m.items = make([]int, 0)
}
