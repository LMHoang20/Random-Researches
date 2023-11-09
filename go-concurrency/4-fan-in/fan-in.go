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
		close(c)
	}()
	return c
}

func shorterFunction() chan string {
	c := make(chan string)
	go func() {
		c <- "I'm done already"
		close(c)
	}()
	return c
}

func fanIn(a chan string, b chan string) {
	c := make(chan string)
	go func() {
		for str := range a {
			c <- str
		}
	}()
	go func() {
		for str := range b {
			c <- str
		}
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func main() {
	fanIn(longerFunction(), shorterFunction())
}
