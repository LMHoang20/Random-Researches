package main

import (
	"fmt"
	"time"
)

func panicFunc(n int) {
	if n == 0 {
		panic("something bad happened")
	}
	fmt.Println("hello", n)
	panicFunc(n - 1)
}

func main() {
	go panicFunc(5)
	go panicFunc(10)

	time.Sleep(2 * time.Second)
}
