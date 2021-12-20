package utils

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadLines(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	return txtlines
}

func ParseLineToIntArray(line, sep string) []int {
	var result []int
	for _, item := range strings.Split(line, ",") {
		if val, err := strconv.Atoi(item); err == nil {
			result = append(result, val)
		} else {
			log.Fatalf("%v, %v\n", "parse int", err)
		}
	}
	return result
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MinArray(x []int) int {
	var min int = math.MaxInt64
	for _, val := range x {
		min = Min(min, val)
	}
	return min
}

func MaxArray(x []int) int {
	var max int = 0
	for _, val := range x {
		max = Max(max, val)
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
