package main

import (
	"log"
	"time"
)

func trace(name string) func() {
	start := time.Now()

	return func() {
		log.Printf("%s: start: %v, end: %v\n", name, start, time.Now())
	}
}

func TestTraceTime() {
	defer trace("TestTraceTime")()
	time.Sleep(2)
}

func main() {
	TestTraceTime()
}
