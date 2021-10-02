package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum int64
	fmt.Println(sum)

	atomic.AddInt64(&sum, 1)
	fmt.Println(sum)

	mu := sync.Mutex{}
	mu.Lock()
	sum += 1
	mu.Unlock()
	fmt.Println(sum)

	var av atomic.Value
	wallace := ninja{name: "Wallace"}
	av.Store(wallace)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		w := av.Load().(ninja)
		w.name = "Not Wallace"
		av.Store(w)
	}()
	wg.Wait()

	fmt.Println(av.Load())
}

type ninja struct {
	name string
}
