package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2025/day1/input.txt")
	lines := strings.Split(string(dat), "\n")
	commands := make([]Command, 0)
	for _, line := range lines {
		command := Command{}
		if line[0] == 'L' {
			command.IsLeft = true
		} else {
			command.IsLeft = false
		}
		num, _ := strconv.Atoi(line[1:])
		command.Num = num
		commands = append(commands, command)
	}
	curPos := 50
	ans := 0
	for _, command := range commands {
		if command.IsLeft {
			curPos = (curPos - command.Num) % 100
		} else {
			curPos = (curPos + command.Num) % 100
		}
		if curPos == 0 {
			ans++
		}
	}
	fmt.Println("Part 1:", ans)

	curPos = 50
	ans = 0
	for _, command := range commands {
		if command.IsLeft {
			for i := 0; i < command.Num; i++ {
				curPos = (curPos - 1) % 100
				if curPos == 0 {
					ans++
				}
			}
		} else {
			for i := 0; i < command.Num; i++ {
				curPos = (curPos + 1) % 100
				if curPos == 0 {
					ans++
				}
			}
		}
	}
	fmt.Println("Part 2:", ans)
}

type Command struct {
	IsLeft bool
	Num    int
}
