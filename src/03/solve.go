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
	fmt.Println("part 1", gammaN*epsilonN)

	var oxygen int64
	pos := 0
	var filtered []string = lines[0:]
	found := false
	for {
		if filtered, found = filterLines(filtered, pos, true); found {
			oxygen, _ = strconv.ParseInt(filtered[0], 2, 64)
			break
		}
		fmt.Println(filtered)
		pos++
	}

	var co2 int64
	pos = 0
	filtered = lines[0:]
	found = false
	for {
		if filtered, found = filterLines(filtered, pos, false); found {
			co2, _ = strconv.ParseInt(filtered[0], 2, 64)
			break
		}
		pos++
	}

	fmt.Println("part 2", oxygen, co2, oxygen*co2)
}

func filterLines(lines []string, pos int, more bool) ([]string, bool) {
	zeros, ones := 0, 0
	var zeroRows, oneRows []string
	for _, line := range lines {
		bit := strings.Split(line, "")[pos]
		if bit == "0" {
			zeros++
			zeroRows = append(zeroRows, line)
		} else {
			ones++
			oneRows = append(oneRows, line)
		}
	}

	fmt.Println(more, pos, zeros, len(zeroRows), ones, len(oneRows))
	if more {
		if ones > zeros || ones == zeros {
			return oneRows, len(oneRows) == 1
		} else {
			return zeroRows, len(zeroRows) == 1
		}
	} else {
		if ones > zeros || ones == zeros {
			return zeroRows, len(zeroRows) == 1
		} else {
			return oneRows, len(oneRows) == 1
		}
	}
}
