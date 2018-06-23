package main

import (
	"fmt"
)

const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	HAHA = 100
)

// iota is zero again
const (
	HEHE = iota
)

func main() {
	fmt.Printf("lsl-debug: %d, %d, %d, %d, %d\n", KB, MB, GB, HAHA, HEHE)
}
