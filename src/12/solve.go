package main

import (
	"adventofcode2021/src/utils"
	"fmt"
)

type Point struct {
	x, y int
}

type Cell struct {
	level   int
	flashed bool
}

func main() {
	lines := utils.ReadLines("test")
	for _, line := range lines {
		fmt.Println(line)
	}
}
