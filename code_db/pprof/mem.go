package main

import (
	"os"
	"runtime/pprof"
)

var memProfPath = "./mem.prof"

func wasteMem() {
	for i := 0; i < 100; i++ {
		_ = make([]byte, 0, 1024*1024)
	}
}

/**
 * go build -o mem mem.go
 * go tool pprof mem mem.prof
 */
func main() {
	f, err := os.Create(memProfPath)
	if err != nil {
		panic("failed to create profile file")
	}
	defer f.Close()

	// function call should be ahead
	wasteMem()

	pprof.WriteHeapProfile(f)
}
