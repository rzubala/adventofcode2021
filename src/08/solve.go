package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"math"
	"math/bits"
	"strings"
)

var uniqueLengts = []int{2, 3, 4, 7}

func main() {

	lines := utils.ReadLines("data")
	cnt1 := 0
	cnt2 := 0
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		output := strings.Split(parts[1], " ")
		for _, seq := range output {
			for _, l := range uniqueLengts {
				if l == len(seq) {
					cnt1++
					break
				}
			}
		}
		cnt2 += part2(parts[0], parts[1])
	}
	fmt.Println("part1", cnt1)
	fmt.Println("part2", cnt2)
}

func toBits(value string) int {
	hash := 0
	for _, b := range strings.Split(value, "") {
		switch b {
		case "a":
			hash |= 1 << 0
		case "b":
			hash |= 1 << 1
		case "c":
			hash |= 1 << 2
		case "d":
			hash |= 1 << 3
		case "e":
			hash |= 1 << 4
		case "f":
			hash |= 1 << 5
		case "g":
			hash |= 1 << 6
		}
	}
	return hash
}

func decode(value string, reference map[int]int) int {
	lv := len(value)
	if lv == 2 {
		return 1
	} else if lv == 3 {
		return 7
	} else if lv == 4 {
		return 4
	} else if lv == 7 {
		return 8
	} else if lv == 5 {
		hash := toBits(value)
		if hash&reference[2] == reference[2] {
			return 3
		} else if bits.OnesCount((uint)(hash&reference[4])) == 3 {
			return 5
		} else {
			return 2
		}
	} else if lv == 6 {
		hash := toBits(value)
		if bits.OnesCount((uint)(hash&reference[4])) == 3 && hash&reference[2] == reference[2] {
			return 0
		} else if bits.OnesCount((uint)(hash&reference[4])) == 3 {
			return 6
		} else {
			return 9
		}
	} else {
		return 0
	}
}

func part2(input string, output string) int {
	inputReferenceSymbols := make(map[int]int)
	for _, i := range strings.Split(input, " ") {
		for _, l := range uniqueLengts {
			if l == len(i) {
				inputReferenceSymbols[l] = toBits(i)
				break
			}
		}
	}

	sum := 0
	for i, o := range strings.Split(output, " ") {
		value := decode(o, inputReferenceSymbols)
		sum += value * int(math.Pow10(3-i))
	}
	return sum
}
