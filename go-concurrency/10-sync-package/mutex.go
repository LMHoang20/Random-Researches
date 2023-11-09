package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func printABunchOfNumbers(withLock bool) {
	if withLock {
		mutex.Lock()
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(10 * time.Millisecond)
	}
	if withLock {
		mutex.Unlock()
	}
}

func main() {
	mutex = sync.Mutex{}

	fmt.Println("Printing with locks")
	go printABunchOfNumbers(true)
	go printABunchOfNumbers(true)

	time.Sleep(1 * time.Second)

	fmt.Println()
	fmt.Println("Printing without locks")
	go printABunchOfNumbers(false)
	go printABunchOfNumbers(false)

	time.Sleep(1 * time.Second)
}
