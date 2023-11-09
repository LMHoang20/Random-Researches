package main

import (
	"math/rand"
	"time"
)

func chanSelector(a chan string, b chan string) {
	for {
		select {
		case str := <-a:
			println(str)
		case str := <-b:
			println(str)
		case <-time.After(1 * time.Second):
			println("timeout")
			return
		}
	}
}

func randomGoroutine(str string) chan string {
	c := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		c <- str
	}()
	return c
}

func main() {
	chanSelector(randomGoroutine("Hello"), randomGoroutine("World"))
}
