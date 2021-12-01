package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strconv"
)

func main() {
	var lines = utils.ReadLines()
	var depths []int
	for _, value := range lines {
		if n, err := strconv.Atoi(value); err == nil {
			depths = append(depths, n)
		}

	}
	calcInc := calcIncreases()
	cnt1, cnt2 := 0, 0
	for _, depth := range depths {
		cnt1 = calcInc(depth)
	}
	fmt.Println("part1", cnt1)
	for i := 3; i < len(depths); i++ {
		win1 := depths[i-1] + depths[i-2] + depths[i-3]
		win2 := depths[i] + depths[i-1] + depths[i-2]
		if win2 > win1 {
			cnt2++
		}
	}
	fmt.Println("part2", cnt2)
}

func calcIncreases() func(int) int {
	prev, cnt := -1, 0
	return func(i int) int {
		if prev > 0 && i > prev {
			cnt++
		}
		prev = i
		return cnt
	}
}
