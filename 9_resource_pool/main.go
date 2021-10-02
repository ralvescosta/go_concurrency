package main

import (
	"fmt"
	"sync"
)

func main() {
	var numMemPieces int

	memPool := &sync.Pool{
		New: func() interface{} {
			numMemPieces++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	const numWokers = 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(numWokers)
	for i := 0; i < numWokers; i++ {
		go func() {
			defer wg.Done()
			mem := memPool.Get().(*[]byte)
			defer memPool.Put(mem)
		}()
	}
	wg.Wait()

	fmt.Printf("%d numMemPueces were created", numMemPieces)
}
