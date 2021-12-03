package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strconv"
)

func main() {
	var lines = utils.ReadLines("data")
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

	calcWindow := memo(func(i int) int {
		return depths[i] + depths[i-1] + depths[i-2]
	})
	for i := 3; i < len(depths); i++ {
		win1 := calcWindow(i - 1)
		win2 := calcWindow(i)
		if win2 > win1 {
			cnt2++
		}
	}
	fmt.Println("part2", cnt2)
}

func calcIncreases() func(int) int { //closure
	prev, cnt := -1, 0
	return func(i int) int {
		if prev > 0 && i > prev {
			cnt++
		}
		prev = i
		return cnt
	}
}

func memo(fn func(int) int) func(int) int { //memoizatoin
	results := make(map[int]int)
	return func(i int) int {
		if value, ok := results[i]; ok {
			return value
		}
		res := fn(i)
		results[i] = res
		return res
	}
}
