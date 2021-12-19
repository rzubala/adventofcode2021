package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.ReadLines("data")
	points := 0
	for _, line := range lines {
		instructions := strings.Split(line, "")
		points += checkCorrupted(instructions)
	}
	fmt.Println("part1", points)
}

func checkCorrupted(instructions []string) int {
	open := make([]string, 0)
	//close := make([]string, 0)
	for _, i := range instructions {
		if isOpen(i) {
			open = append(open, i)
		} else {
			var lastOpen string
			lastOpen, open = open[len(open)-1], open[:len(open)-1]
			if !isMatch(lastOpen, i) {
				return getPoints(i)
			}
		}
	}
	return 0
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

func isClose(ins string) bool {
	return ins == ")" || ins == "]" || ins == "}" || ins == ">"
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
