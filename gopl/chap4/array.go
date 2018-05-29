package main

import (
	"fmt"
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

func main() {
	array_basic()

	array_initialize()
}
