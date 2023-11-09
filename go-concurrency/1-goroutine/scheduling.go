package main

import (
	"fmt"
	"time"
)

func someGoroutine() {
	for i := 0; i < 100; i++ {
		fmt.Println("hello from goroutine", i)
	}
}

func someOtherGoroutine() {
	for i := 100; i < 200; i++ {
		fmt.Println("hello from other goroutine", i)
	}
}

func main() {
	go someGoroutine()
	go someOtherGoroutine()

	fmt.Println("main function")

	time.Sleep(5 * time.Second)
}
