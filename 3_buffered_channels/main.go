package main

import (
	"fmt"
)

func main() {
	bufferedChannel()
}

func deadLockInSimpleChannel() {
	ch := make(chan string)
	ch <- "Hello world"
	fmt.Println(<-ch)
}

func bufferedChannel() {
	ch := make(chan string, 1)
	ch <- "Hello world"
	fmt.Println(<-ch)
}
