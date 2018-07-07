package main

import (
	"fmt"
	"strconv"
)

func TestParseInt() {
	strs := []string{
		"-12345",
		"12345",
		"1234,",
		",1234",
		" 1234",
		"0x12",
		"010",
	}

	for _, s := range strs {
		i, err := strconv.ParseInt(s, 10, 64)
		fmt.Printf("[% 6s] %d, %v\n", s, i, err)
	}
}

func main() {
	TestParseInt()
}
