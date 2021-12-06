package main

import (
	"adventofcode2021/src/utils"
	"fmt"
	"strconv"
	"strings"
)

var maxAge = 9

func main() {
	var countPerAge = make([]int, maxAge)
	lines := utils.ReadLines("data")
	for _, item := range strings.Split(lines[0], ",") {
		if num, err := strconv.Atoi(item); err == nil {
			countPerAge[num]++
		}
	}
	count(countPerAge, 80)
	count(countPerAge, 256)
}

func count(countPerAge []int, days int) {
	for day := 0; day < days; day++ {
		var nextDayCount = make([]int, maxAge)
		for i := range countPerAge {
			dst := (i + 1) % len(countPerAge)
			cnt := countPerAge[dst]
			switch i {
			case 6:
				cnt += countPerAge[0]
			case 8:
				cnt = countPerAge[0]
			}
			nextDayCount[i] = cnt
		}
		countPerAge = nextDayCount
	}
	sum := 0
	for _, item := range countPerAge {
		sum += item
	}
	fmt.Println(days, sum)
}
