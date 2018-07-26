package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}

		close(naturals)
	}()

	go func() {
		/*
		 *         for {
		 *             x, ok := <-naturals
		 *
		 *             if ok {
		 *                 squares <- (x * x)
		 *
		 *             } else {
		 *                 close(squares)
		 *                 break
		 *
		 *             }
		 *         }
		 */
		for x := range naturals {
			squares <- (x * x)
		}

		close(squares)
	}()

	/*
	 *     for {
	 *         v, ok := <-squares
	 *         if ok {
	 *
	 *             fmt.Println(v)
	 *         } else {
	 *
	 *             break
	 *         }
	 *     }
	 */
	for v := range squares {
		fmt.Println(v)
	}
}
