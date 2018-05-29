package main

import (
	"fmt"
)

const (
	_  = 1 << (10 * iota)
	KB // 1 << (10 * 1)
	MB // 1 << (10 * 2)
	GB
	TB
	PB
	EB
	ZB
	YB // 1 << (10 * 8)
)

func main() {
}
