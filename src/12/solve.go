package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strings"
)

func main() {
	nodes := make(map[string][]string)
	lines := utils.ReadLines("test")
	for _, line := range lines {
		parts := strings.Split(line, "-")

		add(parts[0], parts[1], nodes)
		add(parts[1], parts[0], nodes)
	}

	for key := range nodes {
		fmt.Println(key, nodes[key])
	}
}

func add(node1, node2 string, nodes map[string][]string) {
	ns1, ok1 := nodes[node1]
	if !ok1 {
		ns1 = make([]string, 0)
	}
	ns1 = append(ns1, node2)
	nodes[node1] = ns1
}
