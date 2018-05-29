package main

import (
	"fmt"
	"sort"
)

func array_basic() {
	var a [3]int
	fmt.Println(a)

	fmt.Println(len(a))
	fmt.Println(a[len(a)-1])

	for index, value := range a {
		fmt.Printf("%d: %d\n", index, value)
	}

	fmt.Println()
	var b [3]int = [3]int{1, 2, 3}
	for index, value := range b {
		fmt.Printf("%d: %d\n", index, value)
	}

	fmt.Println()
	c := [3]int{1, 2}
	for index, value := range c {
		fmt.Printf("%d: %d\n", index, value)
	}
}

func array_initialize() {
	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(symbol)

	int_array := [...]int{1: 10, 2: 20}
	fmt.Println(int_array)
}

func slice_basic() {
	s := []int{}
	fmt.Printf("[s] len: %d, cap: %d\n", len(s), cap(s))
	// s[0] = 1 // panic

	s = []int{1, 2, 3}
	fmt.Printf("[s] len: %d, cap: %d\n", len(s), cap(s))

	var s1 []int
	fmt.Printf("[s1] len: %d, cap: %d\n", len(s1), cap(s1))
	// s1[0] = 1 // panic

	s2 := make([]int, 0, 10)
	fmt.Printf("[s2] len: %d, cap: %d\n", len(s2), cap(s2))

	s2 = make([]int, 1, 10)
	fmt.Printf("[s2] len: %d, cap: %d\n", len(s2), cap(s2))
	// fmt.Println(s2[1]) // panic
}

func slice_op() {
	var rSlice []rune
	fmt.Printf("rSlice: len: %d, cap: %d\n", len(rSlice), cap(rSlice))

	for _, r := range "Hello 世界" {
		rSlice = append(rSlice, r)
	}

	fmt.Printf("%q\n", rSlice)
	fmt.Printf("rSlice: len: %d, cap: %d\n", len(rSlice), cap(rSlice))
}

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

type Employee struct {
	ID     int
	Name   string
	salary int
}

func struct_basic() {
	boss := Employee{1, "Bob", 1000}
	fmt.Println(boss)

	boss = Employee{ID: 1, Name: "bob", salary: 1000}
	fmt.Println(boss)

	fmt.Printf("%v\n", boss)
	fmt.Printf("%#v\n", boss)
}

func main() {
	/*
	 * array_basic()
	 */

	/*
	 * array_initialize()
	 */

	/*
	 * slice_basic()
	 */

	/*
	 * slice_op()
	 */

	/*
	 * map_basic()
	 */

	/*
	 * map_op()
	 */
	struct_basic()
}
