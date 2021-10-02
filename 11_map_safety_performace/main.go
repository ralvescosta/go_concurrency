package main

import (
	"fmt"
	"sync"
)

func main() {
	syncMap := sync.Map{}
	syncMap.Store(1, 1)
	syncMap.Range(
		func(key, value interface{}) bool {
			fmt.Printf("%d - %d", key, value)
			return true
		})
}
