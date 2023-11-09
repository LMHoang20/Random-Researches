package main

import "time"

func fanOut(receiver chan string, worker1 chan string, worker2 chan string) {
	for str := range receiver {
		select {
		case worker1 <- str:
		case worker2 <- str:
		}
	}
	close(worker1)
	close(worker2)
}

func worker(num int) chan string {
	c := make(chan string)
	go func() {
		for str := range c {
			println("worker", num, "receive", str)
		}
	}()
	return c
}

func main() {
	receiver := make(chan string)
	worker1 := worker(1)
	worker2 := worker(2)
	go fanOut(receiver, worker1, worker2)
	for i := 0; i < 10; i++ {
		receiver <- "Hello"
	}
	close(receiver)

	time.Sleep(1 * time.Second)
}
