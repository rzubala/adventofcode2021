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
	lines := utils.ReadLines("data")

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

	var bingos = make(map[int]bool)
	for _, number := range numbers {
		for b, board := range boards {
			if _, ok := bingos[b]; ok {
				continue
			}
			for y, row := range board {
				if _, ok := bingos[b]; ok {
					break
				}
				for x, p := range row {
					if p.value == number {
						p.Mark()
						board[y][x] = p
					}
					if checkBingo(board, x, y) {
						fmt.Println("Bingo", b, number, number*sumUnmarked((board)))
						bingos[b] = true
						break
					}
				}
			}
		}
	}
}

func (p *Point) Mark() {
	p.mark = true
}

func sumUnmarked(board [][]Point) int {
	sum := 0
	for _, row := range board {
		for _, p := range row {
			if !p.mark {
				sum += p.value
			}
		}
	}
	return sum
}

func checkBingo(board [][]Point, x, y int) bool {
	bingo := true
	for _, p := range board[y] {
		if !p.mark {
			bingo = false
			break
		}
	}
	if bingo {
		return true
	}
	bingo = true
	for _, row := range board {
		if !row[x].mark {
			bingo = false
			break
		}
	}
	return bingo
}

// func printBoard(boards [][][]Point, pos int) {
// 	fmt.Println("***", pos, "***")
// 	for y := 0; y < sizeY; y++ {
// 		lineStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(boards[pos][y])), " "), "[]")
// 		fmt.Println(lineStr)
// 	}
// }
