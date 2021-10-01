package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var missionCompleted bool

func main() {
	var wg sync.WaitGroup
	wg.Add(100)

	var once sync.Once

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			if foundTreasure() {
				once.Do(markMissionCompleted)
			}
		}()
	}
	wg.Wait()

	checkMissionCompletion()
}

func checkMissionCompletion() {
	if missionCompleted {
		fmt.Println("Mission is now completed.")
	} else {
		fmt.Println("Mission was a failure.")
	}
}

func markMissionCompleted() {
	fmt.Println("[::markMissionCompleted]")
	missionCompleted = true
}

func foundTreasure() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
