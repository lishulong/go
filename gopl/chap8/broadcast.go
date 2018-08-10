package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(idx int, wg *sync.WaitGroup, done chan struct{}) {
	defer wg.Done()

	_, ok := <-done
	if !ok {
		fmt.Printf("done channel is closed, quit %d\n", idx)
	}
}

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})

	for idx := 0; idx < 10; idx++ {
		wg.Add(1)
		go worker(idx, &wg, done)
	}

	time.Sleep(10 * time.Millisecond)

	close(done)
	wg.Wait()

	fmt.Printf("see you!\n")
}
