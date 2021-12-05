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

func main() {
	var points1 = make(map[Point]int)
	var points2 = make(map[Point]int)
	lines := utils.ReadLines("data")
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		from := strings.Split(parts[0], ",")
		to := strings.Split(parts[1], ",")
		x1, _ := strconv.Atoi(from[0])
		y1, _ := strconv.Atoi(from[1])
		x2, _ := strconv.Atoi(to[0])
		y2, _ := strconv.Atoi(to[1])
		if x1 == x2 || y1 == y2 {
			for x := utils.Min(x1, x2); x <= utils.Max(x1, x2); x++ {
				for y := utils.Min(y1, y2); y <= utils.Max(y1, y2); y++ {
					p := Point{x: x, y: y}
					inc(points1, p)
					inc(points2, p)
				}
			}
		} else {
			addDiagonalPoints(points2, x1, y1, x2, y2)
		}
	}
	fmt.Println("part1", sum(points1))
	fmt.Println("part2", sum(points2))
}

func addDiagonalPoints(points map[Point]int, x1, y1, x2, y2 int) {
	dx := utils.Abs(x1-x2) + 1
	x := x1
	y := y1
	for ix := 0; ix < dx; ix++ {
		p := Point{x: x, y: y}
		inc(points, p)
		if x1 > x2 {
			x--
		} else if x1 < x2 {
			x++
		}
		if y1 > y2 {
			y--
		} else if y1 < y2 {
			y++
		}
	}
}

func inc(points map[Point]int, key Point) {
	cnt, ok := points[key]
	if ok {
		cnt++
	} else {
		cnt = 1
	}
	points[key] = cnt
}

func sum(points map[Point]int) int {
	res := 0
	for _, cnt := range points {
		if cnt > 1 {
			res++
		}
	}
	return res
}
