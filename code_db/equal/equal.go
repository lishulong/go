package main

import (
	"fmt"
)

func equalStruct() {
	type Person struct {
		Name string
		Age  int
	}

	p1 := &Person{"a", 1}
	p2 := &Person{"a", 1}

	fmt.Printf("p1 == p2 ? %t\n", p1 == p2)
	fmt.Printf("*p1 == *p2 ? %t\n", *p1 == *p2)

	var sli1 []string
	fmt.Printf("sli: %v\n", sli1)
	sli1 = nil
	sli1 = append(sli1, "hello")
	fmt.Printf("sli: %v\n", sli1)
}

func main() {
	equalStruct()
}
