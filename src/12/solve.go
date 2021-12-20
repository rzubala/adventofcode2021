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
	fmt.Println("Part1", findPaths(nodes, false))
	fmt.Println("Part2", findPaths(nodes, true))
}

func findPaths(nodes map[string][]string, part2 bool) int {
	pathsStack := make([][]string, 0)
	start := make([]string, 0)
	start = append(start, "start")
	pathsStack = append(pathsStack, start)
	paths := 0
	pathHashes := make([]string, 0)

	for len(pathsStack) > 0 {
		var currentPath []string
		currentPath, pathsStack = pathsStack[len(pathsStack)-1], pathsStack[:len(pathsStack)-1]
		node := currentPath[len(currentPath)-1]
		if node == "end" {
			hash := strings.Join(currentPath, "")
			newPath := true
			for _, tmp := range pathHashes {
				if tmp == hash {
					newPath = false
					break
				}
			}
			if newPath {
				paths++
			}
			continue
		}

		nextNodes := nodes[node]
		for _, nn := range nextNodes {
			visited := 0
			addPath := true
			if nn == "start" {
				continue
			}
			if utils.IsLower(nn) {
				for _, tmpn := range currentPath {
					if tmpn == nn {
						if !part2 {
							addPath = false
							break
						} else {
							if visited == 1 {
								addPath = false
								break
							}
							visited++
						}
					}
				}
			}
			if addPath {
				nextPath := make([]string, 0)
				nextPath = append(nextPath, currentPath...)
				nextPath = append(nextPath, nn)
				pathsStack = append(pathsStack, nextPath)
			}
		}
	}
	return paths
}

func add(node1, node2 string, nodes map[string][]string) {
	ns1, ok1 := nodes[node1]
	if !ok1 {
		ns1 = make([]string, 0)
	}
	ns1 = append(ns1, node2)
	nodes[node1] = ns1
}
