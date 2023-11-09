package main

import "fmt"

func prepareChan(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	// without closing the channel,
	// the range loop will block
}

func main() {
	c := make(chan int)
	go prepareChan(c)

	for i := range c {
		fmt.Println(i)
	}

	// receive after close will not panic
	i, more := <-c
	fmt.Println(i, more)

	// send after close will panic
	c <- 1
}
