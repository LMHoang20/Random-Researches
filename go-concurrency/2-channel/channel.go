package main

import (
	"fmt"
	"time"
)

func receiveFromChan(c <-chan string) string {
	message := <-c
	return message
}

func sendToChan(c chan<- string, message string) {
	c <- message
}

func main() {
	c := make(chan string)
	go sendToChan(c, "Hello, World!")

	message := receiveFromChan(c)

	fmt.Println(message)

	time.Sleep(1 * time.Second)
}
