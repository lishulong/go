package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Is inf ?  [%v]\n", math.IsInf(math.Inf(1), 1))
	fmt.Printf("Is -inf ? [%v]\n", math.IsInf(math.Inf(-1), -1))
	fmt.Printf("Is NaN? [%v]\n", math.IsNaN(math.NaN()))
	// NaN != NaN
	fmt.Printf("NaN == NaN ? [%v]\n", math.NaN() == math.NaN())
}
