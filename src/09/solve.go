package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Basin struct {
	h, b int
}

func main() {
	matrix := make([][]Basin, 0)

	for _, line := range utils.ReadLines("data") {
		row := make([]Basin, 0)
		for _, hStr := range strings.Split(line, "") {
			h, err := strconv.Atoi(hStr)
			if err != nil {
				panic("can not parse " + hStr)
			}
			row = append(row, Basin{h, -1})
		}
		matrix = append(matrix, row)
	}

	risk := 0
	lowPoints := make([]Point, 0)
	cnt := 0
	for y, row := range matrix {
		for x, val := range row {
			neighbours := getNeighbours(matrix, x, y)
			nvalues := getValues(matrix, neighbours)
			if isMin(val.h, nvalues) {
				risk += 1 + val.h
				lowPoints = append(lowPoints, Point{x, y})
				matrix[y][x] = Basin{val.h, cnt}
				cnt++
			}
		}
	}
	fmt.Println("part1", risk)

	markBasins(matrix, lowPoints)
}

func markBasins(matrix [][]Basin, lowPoints []Point) {
	stack := make([]Point, 0)
	stack = append(stack, lowPoints...)

	var p Point
	for len(stack) > 0 {
		p, stack = stack[len(stack)-1], stack[:len(stack)-1]
		pv := matrix[p.y][p.x]

		neighbours := getNeighbours(matrix, p.x, p.y)
		for _, np := range neighbours {
			n := matrix[np.y][np.x]
			if n.b < 0 && n.h < 9 {
				stack = append(stack, np)
				matrix[np.y][np.x] = Basin{n.h, pv.b}
			}
		}
	}
	getBasinSizes(matrix)
}

func getBasinSizes(matrix [][]Basin) {
	sizes := make(map[int]int)
	for _, row := range matrix {
		for _, v := range row {
			if v.b > -1 {
				cnt, ok := sizes[v.b]
				if ok {
					cnt++
				} else {
					cnt = 1
				}
				sizes[v.b] = cnt
			}
		}
	}
	values := make([]int, 0, len(sizes))
	for k := range sizes {
		values = append(values, sizes[k])
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	fmt.Println("part2", values[0]*values[1]*values[2])
}

func isMin(h int, nvalues []int) bool {
	for _, v := range nvalues {
		if h >= v {
			return false
		}
	}
	return true
}

func getValues(matrix [][]Basin, neighbours []Point) []int {
	result := make([]int, 0)
	for _, p := range neighbours {
		result = append(result, matrix[p.y][p.x].h)
	}
	return result
}

func getNeighbours(matrix [][]Basin, x, y int) []Point {
	sizex := len(matrix[0])
	sizey := len(matrix)
	result := make([]Point, 0)
	y1 := y
	x1 := x - 1
	if x1 >= 0 {
		result = append(result, Point{x1, y1})
	}
	x1 = x + 1
	if x1 < sizex {
		result = append(result, Point{x1, y1})
	}
	x1 = x
	y1 = y - 1
	if y1 >= 0 {
		result = append(result, Point{x1, y1})
	}
	y1 = y + 1
	if y1 < sizey {
		result = append(result, Point{x1, y1})
	}
	return result
}
