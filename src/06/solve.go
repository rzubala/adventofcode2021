package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("data")
	fishes := make([]int, 0)
	for _, item := range strings.Split(lines[0], ",") {
		if num, err := strconv.Atoi(item); err == nil {
			fishes = append(fishes, num)
		}

	}
	days := 80
	for day := 0; day <= days; day++ {
		fmt.Println(day, len(fishes))
		toAdd := 0
		for i, f := range fishes {
			var nf int
			switch f {
			case 0:
				nf = 6
				toAdd++
			default:
				nf = f - 1
			}
			fishes[i] = nf
		}
		for a := 0; a < toAdd; a++ {
			fishes = append(fishes, 8)
		}
	}
}
