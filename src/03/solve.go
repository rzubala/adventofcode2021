package main

import (
	"adventofcode2021/src/utils"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("data")

	ones := make(map[int]int)
	zeros := make(map[int]int)
	for _, line := range lines {
		for y, bit := range strings.Split(line, "") {
			if bit == "0" {
				elem := zeros[y]
				zeros[y] = elem + 1
			} else {
				elem := ones[y]
				ones[y] = elem + 1
			}
		}
	}

	var gamma bytes.Buffer
	var epsilon bytes.Buffer
	for i := 0; i < len(ones); i++ {
		if ones[i] > zeros[i] {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		} else {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		}
	}
	gammaN, _ := strconv.ParseInt(gamma.String(), 2, 64)
	epsilonN, _ := strconv.ParseInt(epsilon.String(), 2, 64)
	fmt.Println(gammaN, epsilonN, gammaN*epsilonN)
}
