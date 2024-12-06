package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("2024/day5/input.txt")
	str := strings.Split(string(dat), "\n\n")
	rulesStr := strings.Split(str[0], "\n")
	updatesStr := strings.Split(str[1], "\n")

	answerOne := 0
	notValidIndexes := make([]int, 0)
	for i, updateStr := range updatesStr {
		secondNums := getSecondNums(rulesStr, updateStr)
		seen := make(map[int]bool)
		isValid := true
		for _, num := range getNums(updateStr) {
			if firstNums, ok := secondNums[num]; ok {
				for firstNum, _ := range firstNums {
					if _, ok := seen[firstNum]; !ok {
						isValid = false
					}
				}
				if !isValid {
					fmt.Println("not valid!")
					notValidIndexes = append(notValidIndexes, i)
					break
				}
			}
			seen[num] = true
		}
		if isValid {
			answerOne += getMiddle(updateStr)
		}
	}
	fmt.Println("Part 1:", answerOne)
}

func getSecondNums(s []string, update string) map[int]map[int]bool {
	// only include in answer rules that apply to current update.
	set := make(map[int]bool)
	for _, num := range getNums(update) {
		set[num] = true
	}
	ans := make(map[int]map[int]bool)
	for _, line := range s {
		numStr := strings.Split(line, "|")
		first, second := getNum(numStr[0]), getNum(numStr[1])
		if ok := set[first]; !ok {
			continue
		}
		if ok := set[second]; !ok {
			continue
		}
		if _, ok := ans[second]; !ok {
			ans[second] = make(map[int]bool)
		}
		ans[second][first] = true
	}
	return ans
}

func getMiddle(s string) int {
	nums := getNums(s)
	if len(nums)%2 != 1 {
		panic("Can't find middle; updates isn't odd.")
	}
	return nums[len(nums)/2]
}

func getNums(s string) []int {
	ans := make([]int, 0)
	for _, num := range strings.Split(s, ",") {
		ans = append(ans, getNum(num))
	}
	return ans
}

func getNum(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
