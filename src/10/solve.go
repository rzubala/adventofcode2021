package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	lines := utils.ReadLines("data")
	corrupted := 0
	scores := make([]int, 0)
	for _, line := range lines {
		instructions := strings.Split(line, "")
		x, score := checkCorrupted(instructions)
		corrupted += x
		if score > 0 {
			scores = append(scores, score)
		}
	}
	fmt.Println("part1", corrupted)
	sort.Ints(scores)
	fmt.Println("part2", scores[len(scores)/2])
}

func checkCorrupted(instructions []string) (int, int) {
	open := make([]string, 0)
	for _, i := range instructions {
		if isOpen(i) {
			open = append(open, i)
		} else {
			var lastOpen string
			lastOpen, open = open[len(open)-1], open[:len(open)-1]
			if !isMatch(lastOpen, i) {
				return getPoints(i), 0
			}
		}
	}
	total := 0
	for len(open) > 0 {
		var lastOpen string
		lastOpen, open = open[len(open)-1], open[:len(open)-1]
		score := getScore(lastOpen)
		total = 5*total + score

	}
	return 0, total
}

func getScore(i string) int {
	switch i {
	case "(":
		return 1
	case "[":
		return 2
	case "{":
		return 3
	case "<":
		return 4
	}
	panic("unknown " + i)
}

func getPoints(i string) int {
	switch i {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}
	panic("unknown " + i)
}

func isOpen(ins string) bool {
	return ins == "(" || ins == "[" || ins == "{" || ins == "<"
}

func isMatch(open string, close string) bool {
	switch open {
	case "(":
		return close == ")"
	case "{":
		return close == "}"
	case "[":
		return close == "]"
	case "<":
		return close == ">"
	}
	panic("unknown open" + open)
}
