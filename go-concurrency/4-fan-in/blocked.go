package main

import (
	"fmt"
	"time"
)

func longerFunction() chan string {
	c := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c <- "I'm finally done"
	}()
	return c
}

func shorterFunction() chan string {
	c := make(chan string)
	go func() {
		c <- "I'm done already"
	}()
	return c
}

func main() {
	// This will block until the first function returns
	// even though the second function is done already
	fmt.Println(<-longerFunction())
	fmt.Println(<-shorterFunction())
}
