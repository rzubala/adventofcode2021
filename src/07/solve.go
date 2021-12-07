package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var positions []int
	lines := utils.ReadLines("data")
	for _, item := range strings.Split(lines[0], ",") {
		if val, err := strconv.Atoi(item); err == nil {
			positions = append(positions, val)
		}
	}
	fmt.Println(positions)

	var minPos int = 1e10
	for _, pos1 := range positions {
		iterSum := 0
		for _, pos2 := range positions {
			iterSum += utils.Abs(pos1 - pos2)
		}
		minPos = utils.Min(minPos, iterSum)
	}
	fmt.Println(minPos)
}
