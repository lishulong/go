package main

import (
	//"errors"
	"fmt"
	//"reflect"
	"sort"
)

type Person struct {
	Age  int
	Name string
}

func (p *Person) String() string {
	return fmt.Sprintf("{ Name: %s, Age: %d }", p.Name, p.Age)
}

func printPersons(ps []*Person) {
	for idx, p := range ps {
		fmt.Printf("[%d] %s\n", idx, p)
	}
	fmt.Println()
}

func main() {
	p1 := &Person{Age: 1, Name: "b"}
	p2 := &Person{Age: 3, Name: "b"}
	p3 := &Person{Age: 2, Name: "c"}
	p4 := &Person{Age: 3, Name: "c"}
	p5 := &Person{Age: 1, Name: "b"}

	var ps []*Person
	ps = append(ps, p5, p4, p3, p2, p1)

	printPersons(ps)

	// sort
	sort.Slice(ps, func(i, j int) bool {
		pi := ps[i]
		pj := ps[j]

		if pi.Name < pj.Name {
			return true

		} else if pi.Name > pj.Name {
			return false

		}

		return pi.Age < pj.Age
	})

	printPersons(ps)
}
