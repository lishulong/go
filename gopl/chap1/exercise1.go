package main

import (
	"fmt"
	"os"
	"strings"
)

func exercise_1_1_2() {
	var s string

	for index, value := range os.Args {
		s += fmt.Sprintf("%d: %s\n", index, value)
	}

	fmt.Print(s)
}

func exercise_1_3() {
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	exercise_1_1_2()

	exercise_1_3()
}
