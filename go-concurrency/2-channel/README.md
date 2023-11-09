# Don't communicate by sharing memory, share memory by communicating.

A **channel** is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine. The value passed to a channel is called a **message**.

A channel variable in **Golang** is a declared by the keyword `chan`, followed by its value's type. 

A value can be send to a channel using the syntax `channel <- value`. A value can be received from a channel using the syntax `value := <- channel`.

When passing into a function, the channel can be an send-only channel (`chan<-`) or a receive-only channel (`<-chan`).

# Buffered channel

A buffered channel is a channel with a buffer. The buffer holds messages that are ready to be received by a channel. The buffer is a queue, so the first message in is the first message out.

```go
c := make(chan int, n)
```

# Unbuffered channel

An unbuffered channel is a channel without a buffer. Messages are passed through the channel immediately. If there is no goroutine waiting to receive the message, the sender will block until there is.

```go
c := make(chan int)
```

# Behaviour

## Send

When a goroutine tries to send a message to a channel, it will block until another goroutine receives the message. If the channel is unbuffered, the sender will block until the receiver receives the message. If the channel is buffered, the sender will block until the message is placed in the buffer.

## Receive

When a goroutine tries to receive a message from a channel, it will block until another goroutine sends a message.

# Closing Channel

Closing a channel is a way to tell the receiver that no more messages will be sent on the channel. It is a way to communicate to the receiver that the work is done.

A channel should **never** be closed by the receiver, only the sender. This is because when a sender tries to send to closed channel, it will panic. A receiver tries to receive from a closed channel will not panic, it will return the zero value of the channel type and a boolean that says the channel is closed.

It is not necessary to close a channel, but it is a good practice to do so. Closing a channel is neccessary to tell the receiver that there is no more messages to be received. If the receiver doesn't know when to stop receiving, it will block forever, especially when using `range`.

# Using `range` on Channel

```go
for i := range c {
    fmt.Println(i)
}
```

You can loop through a channel using `range`, it will receive messages until the channel is closed, or block if the channel doesn't have any messages. The loop will exit if the boolean value returned by the channel is false, just like intended.


```go

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
}

```