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
	matrix := make([][]int, 0)

	for _, line := range utils.ReadLines("data") {
		row := make([]int, 0)
		for _, hStr := range strings.Split(line, "") {
			h, err := strconv.Atoi(hStr)
			if err != nil {
				panic("can not parse " + hStr)
			}
			row = append(row, h)
		}
		matrix = append(matrix, row)
	}

	risk := 0
	for y, row := range matrix {
		for x, h := range row {
			neighbours := getNeighbours(matrix, x, y)
			nvalues := getValues(matrix, neighbours)
			if isMin(h, nvalues) {
				risk += 1 + h
			}
		}
	}

	print(matrix)
	fmt.Println(risk)
}

func isMin(h int, nvalues []int) bool {
	for _, v := range nvalues {
		if h >= v {
			return false
		}
	}
	return true
}

func getValues(matrix [][]int, neighbours []Point) []int {
	result := make([]int, 0)
	for _, p := range neighbours {
		result = append(result, matrix[p.y][p.x])
	}
	return result
}

func getNeighbours(matrix [][]int, x, y int) []Point {
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

func print(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
