package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	salary int
}

func structBasic() {
	boss := Employee{1, "Bob", 1000}
	fmt.Println(boss)

	boss = Employee{
		ID:     1,
		Name:   "bob",
		salary: 1000,
	}
	fmt.Println(boss)

	fmt.Printf("%v\n", boss)
	fmt.Printf("%#v\n", boss)

	boss2 := Employee{ID: 1, Name: "bob", salary: 1000}
	fmt.Println("boss == boss2? ", boss == boss2)
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius float64
}

func structFields() {
	c := Circle{Point{1, 1}, 2}
	fmt.Printf("%#v\n", c)
}

func main() {
	structBasic()

	structFields()
}
