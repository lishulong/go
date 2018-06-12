package main

import (
	"fmt"
	"os"
)

func TestStat() {
	NEFN := "not_exist"

	if _, err := os.Stat(NEFN); os.IsNotExist(err) {
		fmt.Printf("[%s] not exist\n", NEFN)
	}

	EFN := "stat.go"
	if _, err := os.Stat(EFN); err == nil {
		fmt.Printf("[%s] exist\n", EFN)

	} else if os.IsNotExist(err) {
		fmt.Printf("[%s] not exist\n", EFN)

	} else {
		fmt.Printf("[%s] error occured: %v\n", err.(*os.PathError))

	}

}

func main() {
	TestStat()
}
