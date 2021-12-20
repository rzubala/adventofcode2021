package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Fold struct {
	vert  bool
	value int
}

func main() {
	dots := make(map[Point]bool)
	folds := make([]Fold, 0)
	lines := utils.ReadLines("data")
	var w, h int
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			x, ex := strconv.Atoi(parts[0])
			if ex != nil {
				panic("can not parse " + parts[0])
			}
			y, ey := strconv.Atoi(parts[1])
			if ey != nil {
				panic("can not parse " + parts[1])
			}
			if x > w {
				w = x
			}
			if y > h {
				h = y
			}
			dots[Point{x, y}] = true
			continue
		}
		parts = strings.Split(line, "=")
		if len(parts) == 2 {
			value, e := strconv.Atoi(parts[1])
			if e != nil {
				panic("can not parse " + parts[1])
			}
			dir := strings.Split(parts[0], "")
			var vert bool
			if dir[len(dir)-1] == "x" {
				vert = false
			} else if dir[len(dir)-1] == "y" {
				vert = true
			} else {
				panic("something went wrong")
			}
			folds = append(folds, Fold{vert, value})
		}
	}
	for i, ins := range folds {
		var process func(x, y, f int) Point
		if ins.vert {
			process = processY
		} else {
			process = processX
		}
		dots = fold(dots, w, h, ins.value, process)
		w, h = getCoord(dots, w, h)
		if i == 0 {
			fmt.Println("part1", count(dots, w, h))
		}
	}
	print(dots, w, h)
}

func getCoord(dots map[Point]bool, w, h int) (int, int) {
	nw, nh := math.MinInt64, math.MinInt64
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			_, ok := dots[Point{x, y}]
			if ok {
				if x > nw {
					nw = x
				}
				if y > nh {
					nh = y
				}
			}
		}
	}
	return nw, nh
}

func count(dots map[Point]bool, w, h int) int {
	sum := 0
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			_, ok := dots[Point{x, y}]
			if ok {
				sum++
			}
		}
	}
	return sum
}

func processY(x, y, f int) Point {
	ny := y
	if y > f {
		ny = 2*f - y
	}
	return Point{x: x, y: ny}
}

func processX(x, y, f int) Point {
	nx := x
	if x > f {
		nx = 2*f - x
	}
	return Point{x: nx, y: y}
}

func fold(dots map[Point]bool, w, h, f int, process func(x, y, f int) Point) map[Point]bool {
	folded := make(map[Point]bool)
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			_, ok := dots[Point{x, y}]
			if !ok {
				continue
			}
			p := process(x, y, f)
			folded[p] = true
		}
	}
	return folded
}

func print(dots map[Point]bool, w, h int) {
	fmt.Println()
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			_, ok := dots[Point{x, y}]
			if ok {
				fmt.Print("o")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
