package main

import (
	"fmt"
	"time"
)

func someGoroutine() {
	fmt.Println("hello from goroutine")
}

func main() {
	go someGoroutine()

	go func() {
		fmt.Println("hello from anonymous function")
	}()

	fmt.Println("main function")

	time.Sleep(1 * time.Second)
}
