package main

import (
	"bytes"
	"fmt"
	"net/http/httputil"
)

func TestChunkedReader() {
	src := "7\r\nMozilla\r\n9\r\nDeveloper\r\n7\r\nNetwork\r\n0\r\n\r\n"
	cReader := httputil.NewChunkedReader(bytes.NewBufferString(src))

	data := make([]byte, 1024)
	n, err := cReader.Read(data)
	fmt.Printf("lsl-debug: %d, %v\n", n, err)
	fmt.Printf("data received: %v\n", string(data))
}

func main() {
	TestChunkedReader()
}
