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
	var lines = utils.ReadLines()

	operations := make(map[string]func(*Point, int))
	operations["forward"] = func(p *Point, v int) {
		p.x += v
	}
	operations["up"] = func(p *Point, v int) {
		p.y -= v
	}
	operations["down"] = func(p *Point, v int) {
		p.y += v
	}

	position := Point{x: 0, y: 0}
	for _, line := range lines {
		cmd := strings.Split(line, " ")
		if steps, err := strconv.Atoi(cmd[1]); err == nil {
			operations[cmd[0]](&position, steps)
		}
	}
	fmt.Printf("Position %d, %d -> %d\n", position.x, position.y, position.x*position.y)
}
