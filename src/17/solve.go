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
	} else if p.vx < 0 {
		p.vx++
	}
	p.vy--
}

func (p *probe) inTarget(t target) bool {
	return p.x >= t.x1 && p.x <= t.x2 && p.y >= t.y1 && p.y <= t.y2
}

func (p *probe) outTarget(t target) bool {
	return p.x > t.x2 || p.y < t.y1
}

func main() {
	//var t = target{20, 30, -10, -5}
	var t = target{137, 171, -98, -73}

	var maxY = 0
	cnt := 0
	for vx := t.x2; vx > 0; vx-- {
		for vy := t.y1; vy <= -t.y1; vy++ {
			var p = probe{vx: vx, vy: vy}
			var tmpY = 0
			for {
				p.move()
				tmpY = utils.Max(p.y, tmpY)
				if p.inTarget(t) {
					maxY = utils.Max(maxY, tmpY)
					cnt++
					break
				}
				if p.outTarget(t) {
					break
				}
			}
		}
	}
	fmt.Println("part 1", maxY)
	fmt.Println("part 2", cnt)
}
