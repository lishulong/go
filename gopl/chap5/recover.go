package main

import "fmt"

type Recoverable struct{}

func NeedRecover(recoverable bool) {
	if recoverable {
		panic(Recoverable{})
	}

	panic("non-recoverable")
}

func doRecover(recoverable bool) (err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
		case Recoverable{}:
			err = fmt.Errorf("recoverable panic: %#v", p)
		default:
			panic(p)
		}
	}()

	NeedRecover(recoverable)

	return nil
}

func TestRecover() {
	err := doRecover(true)
	fmt.Println(err)

	doRecover(false)
}

func main() {
	TestRecover()
}
