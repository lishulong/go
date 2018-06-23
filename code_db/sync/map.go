package main

import (
	"fmt"
	"sync"
)

func TestDelDuringRange() {
	smap := new(sync.Map)
	keys := []int{0, 1, 2, 3, 4}

	for idx, val := range keys {
		smap.Store(idx, val)
	}

	cnt := 0
	smap.Range(func(k, v interface{}) bool {
		key := k.(int)
		val := v.(int)

		fmt.Printf("visited: %v, %v\n", key, val)

		if cnt == 0 {
			rmKey := (key + 1) % 5
			smap.Delete(rmKey)
			fmt.Printf("removed: %v\n", rmKey)
		}
		cnt++

		if cnt == 5 {
			return false
		}

		return true
	})

	fmt.Printf("sync.map: len: %v\n", smap)
}

func main() {
	TestDelDuringRange()
}
