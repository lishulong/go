package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http/httputil"
	"strings"
)

var _ = strings.Compare

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
		"Mozallia", "Developer", "Network", "", "",
	}

	data := new(bytes.Buffer)
	chunkedWr := httputil.NewChunkedWriter(data)

	// send one by one
	for _, s := range src {
		n, err := chunkedWr.Write([]byte(s + "\n"))
		fmt.Printf("lsl-debug: %s, %d, %v\n", s, n, err)
	}
	chunkedWr.Close()
	fmt.Printf("data sent: [%v]\n", data.Bytes())
	fmt.Printf("================\n")

	// send all together
	//toSend := strings.Join(src, "\n")
	//fmt.Printf("toSend: [%s]\n", toSend)
	//n, err := chunkedWr.Write([]byte(toSend))
	//chunkedWr.Close()

	//fmt.Printf("send: [%s], %d, %v\n", data.String(), n, err)
	//fmt.Printf("lsl-debug: %s\n", string(data.Bytes()))

	/*
	 *     // read from chunked reader
	 *     chunkedRd := httputil.NewChunkedReader(data)
	 *     buf := make([]byte, 100)
	 *     n, err := chunkedRd.Read(buf)
	 *     fmt.Printf("lsl-debug: %v, %v\n", n, err)
	 *
	 */
	chunkedRd := httputil.NewChunkedReader(data)
	sc := bufio.NewScanner(chunkedRd)

	for cnt := 0; sc.Scan(); cnt++ {
		content := string(sc.Bytes())
		fmt.Printf("received: %d [%s]\n", cnt, content)
	}

	err := sc.Err()
	fmt.Printf("lsl-debug: err: %v\n", err)
}

func main() {
	//TestScaner()

	TestChunkedEncodingScaner()
}
