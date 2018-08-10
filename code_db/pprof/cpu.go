package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

var cpuProfPath = "./cpu.prof"

func SlowFib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return (SlowFib(n-2) + SlowFib(n-1))
}

func wastTime() {
	n := 40
	fmt.Printf("Fib(%d): %d\n", n, SlowFib(n))
}

/**
 * go build -o cpu cpu.go
 * go tool pprof cpu cpu.prof
 */
func main() {
	f, err := os.Create(cpuProfPath)
	if err != nil {
		panic("failed to create profile file")
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	wastTime()
}
