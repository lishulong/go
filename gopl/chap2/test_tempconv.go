package main

import "fmt"

import "go/gopl/chap2/tempconv"

func main() {
	var c tempconv.Celsius = 100
	var f tempconv.Fahrenheit = 100

	fmt.Println(tempconv.CToF(c))
	fmt.Println(tempconv.FToC(f))
}
