package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type testStruct struct {
	a uint64
	b int32
	c [16]byte
}

func basicTypeSize() {
	var i32 int32
	fmt.Printf("unsafe size of int32: %d\n", unsafe.Sizeof(i32))
	fmt.Printf("reflect size of int32: %d\n", reflect.TypeOf(i32).Size())

	var ts testStruct
	fmt.Printf("unsafe size of int32: %d\n", unsafe.Sizeof(ts))
	fmt.Printf("reflect size of int32: %d\n", reflect.TypeOf(ts).Size())
}

func main() {
	basicTypeSize()
}
