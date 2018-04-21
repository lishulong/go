package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for idx, arg := range os.Args[1:] {
		fmt.Println("index: ", idx)
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
