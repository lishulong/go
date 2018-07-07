package main

import (
	"fmt"
	"net/url"
	"reflect"
)

func main() {
	v := url.Values{}
	v.Set("name", "Ava")

	// []string
	fmt.Println(v["name"])
	fmt.Println(reflect.TypeOf(v["name"]))

	// string
	fmt.Println(v.Get("name"))
	fmt.Println(reflect.TypeOf(v.Get("name")))

	// append
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	fmt.Println(v["friend"])
	// Get returns the first one
	fmt.Println(v.Get("friend"))
	fmt.Println(v.Get("friend"))
	fmt.Println(reflect.TypeOf(v.Get("friend")))

	fmt.Printf("%#v\n", v)
}
