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

	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		instructions[parts[0]] = parts[1]
	}

	//fmt.Println(polymer)
	//fmt.Println(instructions)

	steps := 10
	for step := 1; step <= steps; step++ {
		tmp := ""
		var last string
		for i := 0; i < len(polymer)-1; i++ {
			pair := polymer[i] + polymer[i+1]
			toInsert, ok := instructions[pair]
			if !ok {
				panic("pair not found " + pair)
			}
			tmp += polymer[i] + toInsert
			last = polymer[i+1]
		}
		tmp += last
		//fmt.Println(step, tmp)
		polymer = strings.Split(tmp, "")
	}

	stats := make(map[string]int)
	for _, n := range polymer {
		cnt, ok := stats[n]
		if ok {
			cnt++
		} else {
			cnt = 1
		}
		stats[n] = cnt
	}
	min, max := math.MaxInt64, math.MinInt64
	for _, v := range stats {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println("part1", max-min)
}
