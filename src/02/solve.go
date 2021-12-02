package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y, aim int
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
	fmt.Printf("Part1, position %d, %d -> %d\n", position.x, position.y, position.x*position.y)

	operations["forward"] = func(p *Point, v int) {
		p.x += v
		p.y += p.aim * v
	}
	operations["up"] = func(p *Point, v int) {
		p.aim -= v
	}
	operations["down"] = func(p *Point, v int) {
		p.aim += v
	}
	position = Point{x: 0, y: 0}
	for _, line := range lines {
		cmd := strings.Split(line, " ")
		if steps, err := strconv.Atoi(cmd[1]); err == nil {
			operations[cmd[0]](&position, steps)
		}
	}
	fmt.Printf("Part2, position %d, %d -> %d\n", position.x, position.y, position.x*position.y)
}
