package main

import (
	"adventofcode2021/src/utils"
	"fmt"
)

type probe struct {
	x, y   int
	vx, vy int
}

type target struct {
	x1, x2, y1, y2 int
}

func (p *probe) move() {
	p.x += p.vx
	p.y += p.vy
	if p.vx > 0 {
		p.vx--
	} else {
		p.vx++
	}
	p.vy--
}

func (p *probe) location() {
	fmt.Println(p.x, p.y)
}

func (p *probe) inTarget(t target) bool {
	return p.x >= t.x1 && p.x <= t.x2 && p.y >= t.y1 && p.y <= t.y2
}

func (p *probe) outTarget(t target) bool {
	return p.x > t.x2 || p.y < t.y2
}

func main() {
	//var t = target{20, 30, -10, -5}
	var t = target{137, 171, -98, -73}

	var maxY = 0
	for vx := t.x1; vx > 1; vx-- {
		for vy := t.y1; vy < -t.y1; vy++ {
			var p = probe{vx: vx, vy: vy}
			var tmpY = 0
			for {
				p.move()
				tmpY = utils.Max(p.y, tmpY)
				if p.inTarget(t) {
					maxY = utils.Max(maxY, tmpY)
					fmt.Println("in max", maxY)
					break
				}
				if p.outTarget(t) {
					//fmt.Println("OUT")
					break
				}
			}
		}
	}
	fmt.Println(maxY)
}
