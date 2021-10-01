package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock   sync.Mutex
	rmLock sync.RWMutex
	count  int
)

func main() {
	// basic()
	readAndWrite()
}

func basic() {
	interations := 1000

	for i := 0; i < interations; i++ {
		go increment()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Resulted count is: ", count)
}

func increment() {
	lock.Lock()
	defer lock.Unlock()
	count++
}

func readAndWrite() {
	go read()
	go write()

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func read() {
	rmLock.RLock()
	defer rmLock.RUnlock()

	fmt.Println("Read locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Read unlocking")
}

func write() {
	rmLock.Lock()
	defer rmLock.Unlock()

	fmt.Println("Write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Write unlocking")
}
