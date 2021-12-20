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
	lines := utils.ReadLines("test")
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
	print(dots, w, h)
}

func print(dots map[Point]bool, w, h int) {
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
