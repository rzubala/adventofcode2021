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
					cnt1, ok1 := points1[p]
					if ok1 {
						points1[p] = cnt1 + 1
					} else {
						points1[p] = 1
					}
					cnt2, ok2 := points2[p]
					if ok2 {
						points2[p] = cnt2 + 1
					} else {
						points2[p] = 1
					}

				}
			}
		} else {
			addDiagonalPoints(points2, x1, y1, x2, y2)
		}
	}
	res1 := 0
	for _, cnt := range points1 {
		if cnt > 1 {
			res1++
		}
	}
	fmt.Println("part1", res1)
	res2 := 0
	for _, cnt := range points2 {
		if cnt > 1 {
			res2++
		}
	}
	fmt.Println("part2", res2)
}

func addDiagonalPoints(points map[Point]int, x1, y1, x2, y2 int) {
	dx := utils.Abs(x1-x2) + 1
	x := x1
	y := y1
	for ix := 0; ix < dx; ix++ {
		p := Point{x: x, y: y}
		cnt, ok := points[p]
		if ok {
			points[p] = cnt + 1
		} else {
			points[p] = 1
		}
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
