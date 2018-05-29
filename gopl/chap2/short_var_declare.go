package main

import (
	"fmt"
)

func main() {
	i := 1
	fmt.Println(0, i)

	i, j := 3, 4

	// short variable declaration, new i in scope of if-statement
	if i := 1; i >= 1 {
		fmt.Println(1, i)
	}
	fmt.Println(2, i)

	if i = 1; i >= 1 {
		fmt.Println(3, i)
	}
	fmt.Println(4, i)

	fmt.Println(2, j)
}
