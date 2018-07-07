package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http/httputil"
	"strings"
)

func TestScaner() {
	var data string

	for line := 0; line < 10; line++ {
		// \r\n is OK too
		data += fmt.Sprintf("this is %d line\n", line)
	}
	//fmt.Printf("lsl-debug:\n%s", data)

	sc := bufio.NewScanner(bytes.NewBufferString(data))
	for sc.Scan() {
		content := sc.Bytes()
		// no \r\n, they have been trimed
		fmt.Printf("%s\n", content)
	}

	if err := sc.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}
}

func TestChunkedEncodingScaner() {
	src := []string{
		"Mozilla", "Developer", "Network",
	}

	data := new(bytes.Buffer)
	chunkedWr := httputil.NewChunkedWriter(data)

	/*
	 * // send one by one
	 * for _, s := range src {
	 *     n, err := chunkedWr.Write([]byte(s + "\n"))
	 *     fmt.Printf("lsl-debug: %s, %d, %v\n", s, n, err)
	 * }
	 * chunkedWr.Close()
	 */

	// send all together
	toSend := strings.Join(src, "\n")
	n, err := chunkedWr.Write([]byte(toSend))
	chunkedWr.Close()

	fmt.Printf("send: [%s], %d, %v\n", toSend, n, err)

	//fmt.Printf("lsl-debug: %s\n", string(data.Bytes()))

	chunkedRd := httputil.NewChunkedReader(data)
	sc := bufio.NewScanner(chunkedRd)

	for sc.Scan() {
		content := string(sc.Bytes())
		fmt.Printf("received: %s\n", content)
	}
}

func main() {
	//TestScaner()

	TestChunkedEncodingScaner()
}
