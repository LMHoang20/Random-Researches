package main

import "fmt"

func fibonacciGenerator() <-chan int {
	c := make(chan int)
	go func() {
		for i, j := 0, 1; ; i, j = j, i+j {
			c <- i + j
		}
	}()
	return c
}

func main() {
	c := fibonacciGenerator()

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
