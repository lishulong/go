package main

import (
	"fmt"
)

func main() {
	a := 15
	b := 31

	fmt.Printf("%#b, %[1]d, %#[2]o, %#[2]x\n", a, b)
	fmt.Printf("%d\n", a)
	fmt.Printf("%#o\n", a)
	fmt.Printf("%#x\n", a)
}
