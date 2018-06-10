package main

import (
	"bytes"
	"fmt"
	"strings"
)

func compare() {
	s1 := "1.0.0"
	s2 := "1.0.1"

	fmt.Printf("%s > %s ? %t\n", s1, s2, s1 > s2)
	fmt.Printf("%s > %s ? %t\n", s2, s1, s2 > s1)
}

func strByteArray() {
	var aBArray [20]byte
	aStr := "1.0.0"

	/*
	 *     buf := bytes.NewBufferString(aStr)
	 *
	 *     //fmt.Println(len(buf.Bytes()))
	 *     //fmt.Println(cap(buf.Bytes()))
	 *     for idx, b := range buf.Bytes() {
	 *         aBArray[idx] = b
	 *     }
	 *
	 */

	for idx, b := range []byte(aStr) {
		aBArray[idx] = b
	}
	fmt.Println(len([]byte(aStr)))
	fmt.Println(aBArray)

	buf := bytes.NewBuffer(aBArray[:])
	str, _ := buf.ReadString(0)
	fmt.Printf("lsl-debug: %d, %s\n", len(str), str)

	var strBuilder strings.Builder
	strBuilder.Write(aBArray[:])
	str = strBuilder.String()
	fmt.Printf("lsl-debug: %d, %s\n", len(str), str)

	delimiter := strings.IndexByte(str, 0)
	fmt.Printf("lsl-debug: %d, %s\n", delimiter, str[:delimiter])

	str = str[:delimiter]
	fmt.Printf("lsl-debug: %d, %s\n", len(str), str)

	// convert byte array to string
	/*
	 *     bStr := string(aBArray[:])
	 *     fmt.Printf("%s\n", bStr)
	 *
	 *     fmt.Printf("lsl-debug: %d\n", len(bStr))
	 */
	slice := []byte{'a', 'b', 'c'}
	fmt.Println(string(slice))
	fmt.Println(string(slice[:1]))
	fmt.Println(string(slice[:0]))
	fmt.Println("" == string(slice[:0]))

	a := []int{}
	fmt.Println(a[:0])
}

func main() {
	// compare()

	strByteArray()
}
