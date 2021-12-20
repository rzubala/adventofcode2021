package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	lines := utils.ReadLines("data")
	polymer := strings.Split(lines[0], "")
	instructions := make(map[string]string)
	pairs := make(map[string]int)

	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i] + polymer[i+1]
		if cnt, ok := pairs[pair]; ok {
			pairs[pair] = cnt + 1
		} else {
			pairs[pair] = 1
		}
	}

	stats := make(map[string]int)
	for _, n := range polymer {
		if cnt, ok := stats[n]; ok {
			stats[n] = cnt + 1
		} else {
			stats[n] = 1
		}
	}

	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		instructions[parts[0]] = parts[1]
	}

	steps := 40
	for step := 1; step <= steps; step++ {
		tmpPairs := make(map[string]int)
		for pair, cnt := range pairs {
			toInsert := instructions[pair]
			pair1 := strings.Split(pair, "")[0] + toInsert
			if cnt1, ok1 := tmpPairs[pair1]; ok1 {
				tmpPairs[pair1] = cnt1 + cnt
			} else {
				tmpPairs[pair1] = cnt
			}

			pair2 := toInsert + strings.Split(pair, "")[1]
			if cnt2, ok2 := tmpPairs[pair2]; ok2 {
				tmpPairs[pair2] = cnt2 + cnt
			} else {
				tmpPairs[pair2] = cnt
			}
			stats[toInsert] = stats[toInsert] + cnt
		}
		pairs = tmpPairs
		if step == 10 {
			fmt.Println("part1", getValue(stats))
		}
	}
	fmt.Println("part2", getValue(stats))
}

func getValue(stats map[string]int) int {
	min, max := math.MaxInt64, math.MinInt64
	for _, v := range stats {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}
