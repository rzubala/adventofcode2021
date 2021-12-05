package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var points = make(map[Point]int)
	lines := utils.ReadLines("data")
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		from := strings.Split(parts[0], ",")
		to := strings.Split(parts[1], ",")
		x1, _ := strconv.Atoi(from[0])
		y1, _ := strconv.Atoi(from[1])
		x2, _ := strconv.Atoi(to[0])
		y2, _ := strconv.Atoi(to[1])
		if x1 != x2 && y1 != y2 {
			continue
		}
		fmt.Println("***", x1, y1, x2, y2)
		for x := min(x1, x2); x <= max(x1, x2); x++ {
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				p := Point{x: x, y: y}
				cnt, ok := points[p]
				if ok {
					fmt.Println(x, y, cnt)
					points[p] = cnt + 1
				} else {
					fmt.Println(x, y, 1)
					points[p] = 1
				}
			}
		}
	}
	res := 0
	for _, cnt := range points {
		if cnt > 1 {
			res++
		}
	}
	fmt.Println(res)
}
