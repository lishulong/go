package main

import (
	"fmt"
	"reflect"
)

func judgeInterfaceNil(val interface{}) {
	if val == nil {
		fmt.Printf("val is nil\n")
	} else {
		fmt.Printf("val is not nil\n")
	}
}

func TestJudgeInterfaceNil() {
	judgeInterfaceNil(nil)

	var s []string

	fmt.Printf("type of s: %v, %v\n", reflect.TypeOf(s), s)

	// zero value is not nil
	judgeInterfaceNil(s)

	var i int
	judgeInterfaceNil(i)
}

func main() {
	TestJudgeInterfaceNil()
}
