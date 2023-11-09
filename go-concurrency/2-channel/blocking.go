package main

import (
	"fmt"
	"time"
)

func a(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("a send to channel")
		fmt.Println(i, "a")
		c <- i
	}
}

func b(c <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("b receive from channel")
		fmt.Println(<-c, "b")
	}
}

func main() {
	c := make(chan int)
	go a(c)
	go b(c)

	fmt.Println("main function")

	time.Sleep(1 * time.Second)
}
