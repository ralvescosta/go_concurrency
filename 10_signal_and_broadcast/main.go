package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	ready bool
)

func main() {
	// title: Golang Concurrency - Signal & Broadcast [syn.Cond]

	//gettingReadyForMission()
	//gettingReadyForMissionWithCond()
	broadcastStartOfMission()
}

func broadcastStartOfMission() {
	m := sync.Mutex{}
	beeper := sync.NewCond(&m)
	var wg sync.WaitGroup
	wg.Add(3)
	standByForMission(func() {
		fmt.Println("Ninja 1 starting mission.")
		wg.Done()
	}, beeper)
	standByForMission(func() {
		fmt.Println("Ninja 2 starting mission.")
		wg.Done()
	}, beeper)
	standByForMission(func() {
		fmt.Println("Ninja 3 starting mission.")
		wg.Done()
	}, beeper)
	beeper.Broadcast()
	wg.Wait()
	fmt.Println("All Ninjas have started their missions")
}

func standByForMission(fn func(), beeper *sync.Cond) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		beeper.L.Lock()
		defer beeper.L.Unlock()
		beeper.Wait()
		fn()
	}()
	wg.Wait()
}

func gettingReadyForMission() {
	defer elapsed("Getting ready")()

	go gettingReady()
	workIntervals := 0
	for !isReady() {
		workIntervals++
	}
	fmt.Printf("We are now ready! After %d work intervals.\n", workIntervals)
}

func gettingReadyForMissionWithCond() {
	defer elapsed("Getting ready")()

	m := sync.Mutex{}
	c := sync.NewCond(&m)

	go gettingReadyWithCond(c)
	workIntervals := 0
	c.L.Lock()
	for !isReady() {
		workIntervals++
		c.Wait()
	}
	c.L.Unlock()
	fmt.Printf("We are now ready! After %d work intervals.\n", workIntervals)
}

func gettingReady() {
	sleep()
	ready = true
}

func gettingReadyWithCond(c *sync.Cond) {
	sleep()
	ready = true
	c.Signal()
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}

func isReady() bool {
	return ready
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
