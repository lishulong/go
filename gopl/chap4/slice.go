package main

import (
	"fmt"
)

func slice_basic() {
	s1 := []int{}
	fmt.Printf("[s1] len: %d, cap: %d\n", len(s1), cap(s1))
	// s[0] = 1 // panic

	s2 := []int{1, 2, 3}
	fmt.Printf("[s2] len: %d, cap: %d\n", len(s2), cap(s2))

	var s3 []int
	fmt.Printf("[s3] len: %d, cap: %d\n", len(s3), cap(s3))
	// s1[0] = 1 // panic

	s4 := make([]int, 0, 10)
	fmt.Printf("[s4] len: %d, cap: %d\n", len(s4), cap(s4))

	s5 := make([]int, 1, 10)
	fmt.Printf("[s5] len: %d, cap: %d\n", len(s5), cap(s5))
	// fmt.Println(s5[1]) // panic
	//s5[1] = 0

	s6 := []int{1, 2, 3}
	s7 := s6[1:3]

	// s6 and s7 share the last two elements of the underlaying array
	s7[0] = 10
	fmt.Println(s6)
	fmt.Println(s7)

	/* slice only can be compared with nil
	 * s8 := []int{1, 2, 3}
	 * fmt.Printf("s6 == s8 ? %t\n", s6 == s8) // panic
	 */
}

func slice_op() {
	var rSlice []rune
	fmt.Printf("rSlice: len: %d, cap: %d\n", len(rSlice), cap(rSlice))

	for _, r := range "Hello 世界" {
		rSlice = append(rSlice, r)
	}

	fmt.Printf("%q\n", rSlice)
	fmt.Printf("rSlice: len: %d, cap: %d\n", len(rSlice), cap(rSlice))

	s1 := []int{1, 2, 3}
	s2 := make([]int, 2, 3)
	num := copy(s2, s1)
	fmt.Println("copied: ", num)
	fmt.Println("s1: ", s1, " , s2: ", s2)
}

func main() {
	slice_basic()

	slice_op()
}
