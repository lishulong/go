package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func retError(mode int) error {
	if mode == 1 {
		return errors.New("self defined")

	} else {
		return &os.PathError{
			Op:   "op",
			Path: "path",
			Err:  errors.New("path error"),
		}
	}

}
func TestConvertType() {
	err := retError(1)
	fmt.Printf("type: %v, value: %v\n", reflect.TypeOf(err), err)

	pErr, ok := err.(*os.PathError)
	fmt.Printf("not path error:[%v] %v\n", ok, pErr)

	err = retError(2)
	fmt.Printf("type: %v, value: %v\n", reflect.TypeOf(err), err)

	pErr, ok = err.(*os.PathError)
	fmt.Printf("path error: [%v] %v\n", ok, pErr)

}

func main() {
	TestConvertType()
}
