package main

import (
	"fmt"
	"sort"
)

func map_basic() {
	var fMap map[string]int
	fMap = make(map[string]int)
	fMap["hello"] = 1

	ssMap := map[string]string{
		"first":  "one",
		"second": "two",
	}

	ssMap = make(map[string]string)
	ssMap["third"] = "three"
}

func map_op() {
	ages := make(map[string]int)
	ages["bob"] = 1

	delete(ages, "mike")
	ages["bob"]++
	ages["mike"] = 2

	for key, value := range ages {
		fmt.Printf("[%s]: %d\n", key, value)
	}

	names := make([]string, 0, len(ages))
	for n := range ages {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("[%s]: %d\n", name, ages[name])
	}
}

func main() {
	map_basic()

	map_op()
}
