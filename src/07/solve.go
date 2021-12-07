package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"math"
)

func main() {

	lines := utils.ReadLines("data")
	var positions []int = utils.ParseLineToIntArray(lines[0], ",")

	var minCost1 int = math.MaxInt64
	for _, pos1 := range positions {
		iterSum := 0
		for _, pos2 := range positions {
			iterSum += utils.Abs(pos1 - pos2)
		}
		minCost1 = utils.Min(minCost1, iterSum)
	}
	fmt.Println("Part1", minCost1)

	var minCost2 int = math.MaxInt64
	for pos1 := utils.MinArray(positions); pos1 <= utils.MaxArray(positions); pos1++ {
		iterSum := 0
		for _, pos2 := range positions {
			iterSum += moveCost(pos1, pos2)
		}
		minCost2 = utils.Min(minCost2, iterSum)
	}
	fmt.Println("Part2", minCost2)
}

func moveCost(from, to int) int {
	if from == to {
		return 0
	}
	cost := 0
	diff := utils.Abs(to - from)
	for i := 0; i < diff; i++ {
		cost += 1 + i
	}
	return cost
}
