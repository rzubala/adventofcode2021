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

type Cost struct {
	node Node
	cost int
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
	print(nodes, w, h)

	start := Node{0, 0}
	end := Node{w - 1, h - 1}
	costs := make(map[Node]int)
	previous := make(map[Node]Node)
	stack := make([]Cost, 0)
	stack = append(stack, Cost{start, 0})

	for len(stack) > 0 {
		var costNode Cost
		costNode, stack = stack[len(stack)-1], stack[:len(stack)-1]
		node := costNode.node
		if node.x == end.x && node.y == end.y {
			continue
		}
		fmt.Println(node, len(stack), len(previous), len(costs))
		cost := costNode.cost
		nextNodes := node.Neighbours(w, h)
		for _, nextNode := range nextNodes {
			risk := nodes[nextNode]
			nextCost, ok := costs[nextNode]
			newCost := risk + cost
			if ok && newCost < nextCost || !ok {
				costs[nextNode] = newCost
				previous[nextNode] = node
				stack = append(stack, Cost{nextNode, newCost})
			}
		}

	}
	fmt.Println("part1", costs[end])
	//printPath(previous, costs, end, start)
}

func printPath(path map[Node]Node, costs map[Node]int, from, to Node) {
	node := from
	totalCost := 0
	for {
		cost := costs[node]
		totalCost += cost
		fmt.Println(node, cost, totalCost)
		var ok bool
		if node, ok = path[node]; !ok {
			break
		} else {
			if node.x == to.x && node.y == to.y {
				break
			}
		}
	}
}

func print(nodes map[Node]int, w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print(nodes[Node{x, y}])
		}
		fmt.Println()
	}
}
