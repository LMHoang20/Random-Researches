package main

import (
	"fmt"
	"time"
)

func fibonacciGenerator(Quit chan bool) chan int {
	c := make(chan int)
	go func() {
		for i, j := 0, 1; ; i, j = j, i+j {
			select {
			case c <- i + j:
			case <-Quit:
				close(c)
				fmt.Println("Quitting")
				return
			}
		}
	}()
	return c
}

func main() {
	Quit := make(chan bool)

	c := fibonacciGenerator(Quit)

	for i := 0; i < 10; i++ {
		println(<-c)
	}

	Quit <- true

	time.Sleep(1 * time.Second)
}
