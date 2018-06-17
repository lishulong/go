package main

import (
	"bytes"
	"fmt"
)

func TestCapLimit() {
	data := make([]byte, 0, 2)
	buf := bytes.NewBuffer(data)

	n, err := buf.Write([]byte{1, 2, 3})

	fmt.Printf("n: %d, err: %v, buf: %v, data: %v\n", n, err, buf.Bytes(), data)

}

func main() {
	TestCapLimit()
}
