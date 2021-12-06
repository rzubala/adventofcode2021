package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var fishCount [9]int
	lines := utils.ReadLines("data")
	for _, item := range strings.Split(lines[0], ",") {
		if num, err := strconv.Atoi(item); err == nil {
			fishCount[num]++
		}

	}

	days := 256
	var nextCount [9]int
	for day := 0; day < days; day++ {
		for i := 0; i < 8; i++ {
			nextCount[i] = fishCount[(i+1)%9]
		}
		nextCount[6] += fishCount[0]
		nextCount[8] = fishCount[0]
		nextCount, fishCount = fishCount, nextCount
	}
	sum := 0
	for _, item := range fishCount {
		sum += item
	}
	fmt.Println(sum)
}
