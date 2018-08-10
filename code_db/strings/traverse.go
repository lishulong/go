package main

import (
	"fmt"
)

func main() {
	s := "abcd"

	for i, b := range s {
		fmt.Printf("%d: %v, %T\n", i, b, b)
	}

	fmt.Printf("===============\n")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: %v, %T\n", i, s[i], s[i])
	}
}
