package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var lines = readFile()
	calcInc := calcIncreases()
	var cnt int
	for _, value := range lines {
		if i, err := strconv.Atoi(value); err == nil {
			cnt = calcInc(i)
		}
	}
	fmt.Println(cnt)
}

func calcIncreases() func(int) int {
	prev, cnt := -1, 0
	return func(i int) int {
		if prev > 0 && i > prev {
			cnt++
		}
		prev = i
		return cnt
	}
}

func readFile() []string {
	file, err := os.Open("data")
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
