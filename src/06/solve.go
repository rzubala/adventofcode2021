package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var offspring [9]int
	lines := utils.ReadLines("data")
	for _, item := range strings.Split(lines[0], ",") {
		if num, err := strconv.Atoi(item); err == nil {
			offspring[num]++
		}

	}

	days := 256
	var nextGeneration [9]int
	for day := 0; day < days; day++ {
		for i := 0; i < 8; i++ {
			nextGeneration[i] = offspring[(i+1)%9]
		}
		nextGeneration[6] += offspring[0]
		nextGeneration[8] = offspring[0]
		nextGeneration, offspring = offspring, nextGeneration
	}
	sum := 0
	for _, item := range offspring {
		sum += item
	}
	fmt.Println(sum)
}
