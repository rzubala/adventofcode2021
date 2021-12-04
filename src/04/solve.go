package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strconv"
	"strings"
)

var sizeY int = 5
var sizeX int = 5

type Point struct {
	value int
	mark  bool
}

func main() {
	lines := utils.ReadLines("test")

	var numbers []int
	var boards [][][]Point
	for _, item := range strings.Split(lines[0], ",") {
		if item, err := strconv.Atoi(item); err == nil {
			numbers = append(numbers, item)
		}
	}

	var board [][]Point
	rowIndex := 0
	for _, line := range lines[1:] {
		if len(line) == 0 {
			board = make([][]Point, sizeY)
			boards = append(boards, board)
			rowIndex = 0
			continue
		}
		boardRow := make([]Point, sizeX)
		board[rowIndex] = boardRow
		colIndex := 0
		for _, item := range strings.Split(line, " ") {
			if item, err := strconv.Atoi(item); err == nil {
				boardRow[colIndex] = Point{value: item}
				colIndex++
			}
		}
		rowIndex++
	}

	for index, _ := range boards {
		printBoard(boards, index)
	}

	for _, number := range numbers {

	}
}

func printBoard(boards [][][]Point, pos int) {
	fmt.Println("***", pos, "***")
	for y := 0; y < sizeY; y++ {
		lineStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(boards[pos][y])), " "), "[]")
		fmt.Println(lineStr)
	}
}
