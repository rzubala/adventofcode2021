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

type Cell struct {
	level   int
	flashed bool
}

func main() {
	board := make(map[Point]Cell)
	points := make([]Point, 0)
	lines := utils.ReadLines("data")
	var w, h int
	h = len(lines)
	for y, line := range lines {
		values := strings.Split(line, "")
		w = len(values)
		for x, v := range values {
			if level, err := strconv.Atoi(v); err == nil {
				p := Point{x, y}
				board[p] = Cell{level: level}
				points = append(points, p)
			} else {
				panic("Can not parse" + v)
			}
		}
	}
	queue := make([]Point, 0)
	flashes := 0
	step := 1
	for {
		queue = append(queue, points...)
		for len(queue) > 0 {
			var p Point
			p, queue = queue[0], queue[1:]
			c := board[p]
			if c.flashed {
				continue
			}
			nlevel := c.level + 1
			nflashed := c.flashed
			if nlevel == 10 {
				nlevel = 0
				nflashed = true
				queue = append(queue, p.Neighbours(board, w, h)...)
			}
			board[p] = Cell{level: nlevel, flashed: nflashed}
		}
		flashesPerStep := resetFlashed(board, w, h)
		flashes += flashesPerStep
		if step == 100 {
			fmt.Println("part1", flashes)
		}
		if flashesPerStep == w*h {
			fmt.Println("part2", step)
			break
		}
		step++
	}
}

func (p *Point) Neighbours(board map[Point]Cell, w, h int) []Point {
	result := make([]Point, 0)
	for x := p.x - 1; x <= p.x+1; x++ {
		if x < 0 || x >= w {
			continue
		}
		for y := p.y - 1; y <= p.y+1; y++ {
			if y < 0 || y >= h {
				continue
			}
			if x == p.x && y == p.y {
				continue
			}
			p := Point{x, y}
			c := board[p]
			if c.flashed {
				continue
			}
			result = append(result, Point{x, y})
		}
	}
	return result
}

func resetFlashed(board map[Point]Cell, w, h int) int {
	count := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := Point{x, y}
			c := board[p]
			board[p] = Cell{level: c.level, flashed: false}
			if c.flashed {
				count++
			}
		}
	}
	return count
}
