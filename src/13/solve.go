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
	dots := make(map[Point]bool)
	lines := utils.ReadLines("data")
	var w, h int
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			fmt.Println(parts[0], parts[1])
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
		}
	}
	//print(dots, w, h)

	dots = foldX(dots, w, h, 655)
	fmt.Println("part1", count(dots, w, h))
	//print(dots, w, h)

	//dots = foldX(dots, w, h, 5)
	//print(dots, w, h)
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

func foldY(dots map[Point]bool, w, h, f int) map[Point]bool {
	folded := make(map[Point]bool)
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			_, ok := dots[Point{x, y}]
			if !ok {
				continue
			}
			ny := y
			if y > f {
				ny = 2*f - y
			}
			folded[Point{x: x, y: ny}] = true
		}
	}
	return folded
}

func foldX(dots map[Point]bool, w, h, f int) map[Point]bool {
	folded := make(map[Point]bool)
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			_, ok := dots[Point{x, y}]
			if !ok {
				continue
			}
			nx := x
			if x > f {
				nx = 2*f - x
			}
			folded[Point{x: nx, y: y}] = true
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
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
