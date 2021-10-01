package main

import (
	"fmt"
	"sync"
)

func main() {
	evilNinjas := []string{"Tommy", "Johnny", "Bobby"}

	var beeper = sync.WaitGroup{}
	beeper.Add(len(evilNinjas))
	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, &beeper)
	}
	beeper.Wait()
	fmt.Println("Mission Completed")
}

func attack(evilNinja string, beeper *sync.WaitGroup) {
	defer beeper.Done()
	fmt.Println("Attacked evil ninja: ", evilNinja)
}
