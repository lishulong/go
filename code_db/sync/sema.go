package main

import (
	"context"
	"golang.org/x/sync/semaphore"
	"log"
)

func run(sema *semaphore.Weighted, cnt int) {
	log.Printf("get %d\n", cnt)

	sema.Release(1)
}

func TestUsage() {
	var ctx = context.TODO()
	var semaWeight = 10
	sema := semaphore.NewWeighted(int64(semaWeight))

	for cnt := 0; cnt < 2*semaWeight; cnt++ {
		if err := sema.Acquire(ctx, 1); err != nil {
			log.Fatalf("error occured during sema.acquire: %v", err)
		}

		go run(sema, cnt)

		log.Printf("sema.acquire returns: %d\n", cnt)
	}
}

func main() {
	TestUsage()
}
