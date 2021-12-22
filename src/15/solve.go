package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	x, y int
}

func (n *Node) Neighbours(w, h int) []Node {
	result := make([]Node, 0)
	tmp := []Node{{n.x - 1, n.y}, {n.x + 1, n.y}, {n.x, n.y - 1}, {n.x, n.y + 1}}
	for _, n := range tmp {
		if n.IsInside(w, h) {
			result = append(result, n)
		}
	}
	return result
}

func (n *Node) IsInside(w, h int) bool {
	return n.x >= 0 && n.x < w && n.y >= 0 && n.y < h
}

func main() {
	solve(false)
	solve(true)
}

func solve(extended bool) {
	nodes := make(map[Node]int)
	lines := utils.ReadLines("data")
	w, h := 0, len(lines)
	for y, line := range lines {
		values := strings.Split(line, "")
		w = len(values)
		for x, riskStr := range values {
			if risk, err := strconv.Atoi(riskStr); err != nil {
				panic("can not parse " + riskStr)
			} else {
				nodes[Node{x, y}] = risk
			}
		}
	}
	if extended {
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				for j := 0; j < 5; j++ {
					for i := 0; i < 5; i++ {
						if i == 0 && j == 0 {
							continue
						}
						val := (nodes[Node{x, y}] + i + j) % 9
						if val == 0 {
							val = 9
						}
						nodes[Node{x + w*i, y + h*j}] = val
					}
				}
			}
		}
		w *= 5
		h *= 5
	}
	start := Node{0, 0}
	end := Node{w - 1, h - 1}
	costs := make(map[Node]int)
	previous := make(map[Node]Node)
	stack := make([]Node, 0)
	stack = append(stack, start)
	costs[start] = 0

	totalPessimisticCost := pessimisticCost(w, h, 9)

	for len(stack) > 0 {
		var node Node
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if node.x == end.x && node.y == end.y {
			continue
		}
		cost := costs[node]
		endCost, hasEnd := costs[end]
		if cost >= totalPessimisticCost || cost > pessimisticCost(node.x, node.y, 9) || pessimisticCost(node.x, node.y, 9)+pessimisticCost(w-node.x, h-node.y, 9) > pessimisticCost(w, h, 9) {
			continue
		}
		if hasEnd && cost >= endCost {
			continue
		}
		nextNodes := node.Neighbours(w, h)
		for _, nextNode := range nextNodes {
			risk := nodes[nextNode]
			nextCost, ok := costs[nextNode]
			newCost := risk + cost
			if ok && newCost < nextCost || !ok {
				costs[nextNode] = newCost
				previous[nextNode] = node
				stack = append(stack, nextNode)
			}
		}

	}
	fmt.Println("result", extended, costs[end])
}

func pessimisticCost(x, y, w int) int {
	return (x + y) * w
}
