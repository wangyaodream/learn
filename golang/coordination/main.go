package main

import (
	"sync"
)

func doSum(count int, val *int, waitGroup *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		*val++
	}
	waitGroup.Done()
}

func main() {
	counter := 0

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	go doSum(5000, &counter, &waitGroup)
	waitGroup.Wait()
	Printfln("Total: %v", counter)
}
