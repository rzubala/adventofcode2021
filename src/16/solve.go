package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	packetHex := "D2FE28"
	value, version, header := getLiteralValue(hexToBin(packetHex))
	fmt.Println(value, version, header)
}

func getLiteralValue(bits []string) (uint64, uint64, uint64) {
	version := binToDec(bits[:3])
	header := binToDec(bits[3:6])
	if header != 4 {
		panic("header not match")
	}
	literal := []string{}
	stop := false
	cnt := 6
	for !stop {
		chunk := bits[cnt+1 : cnt+5]
		stop = bits[cnt] == "0"
		cnt += 5
		literal = append(literal, chunk...)
	}
	return binToDec(literal), version, header
}

func binToDec(val []string) uint64 {
	num, err := strconv.ParseUint(strings.Join(val, ""), 2, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func hexToBin(hex string) []string {
	val, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		panic("can not parse " + hex)
	}
	bits := []string{}
	for i := 0; i < 24; i++ {
		bit := val & 0x1
		bitStr := strconv.FormatUint(bit, 2)
		bits = append([]string{bitStr}, bits...)
		val = val >> 1
	}
	return bits
}
