package main

import (
	"errors"
	"fmt"
)

type CallBack func(a int) error

func call(cb CallBack) error {
	if cb != nil {
		if err := cb(1); err != nil {
			return err
		}
	}

	fmt.Println("call is done")

	return nil
}

func main() {
	b := 10
	err := call(func(a int) error {
		b += a

		return nil
	})

	fmt.Printf("b is %d, %v\n", b, err)

	err = call(func(a int) error {
		b += a

		return errors.New("hahaha")
	})

	fmt.Printf("b is %d, %v\n", b, err)
}
