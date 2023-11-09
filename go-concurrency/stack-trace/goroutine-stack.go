package main

import (
	"fmt"
	"time"
)

func panicFunc(n int) {
	if n == 0 {
		panic("something bad happened")
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in panicFunc", r, n)
		}
	}()
	fmt.Println(n)
	if n%2 == 1 {
		go panicFunc(n - 1)
	} else {
		panicFunc(n - 1)
	}
}

func main() {
	go panicFunc(10)

	time.Sleep(2 * time.Second)
}
