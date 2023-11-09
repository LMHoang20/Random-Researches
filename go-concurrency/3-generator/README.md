# Function Returns Channel

A Function that returns a channel acts as a layer of abstraction for the underlying goroutines. The receiver doesn't need to know how the messages are generated, it only needs to know how to receive them.

# Block Function Execution

One common usage of this pattern is to block a function execution for a certain ammount of time. Here, `time.After(...)` returns a channel. Trying to receive from this channel will block the execution. After a certain ammount of time, the underlying goroutine will send a message to the channel, unblocking the caller. 

```go
<-time.After(1 * time.Second):
```

# Promise 

This pattern can also be used to implement a Promise-like functionality. The caller can receive from the returned channel to wait for the result of the asynchronous operation. 

Javascript's `async`, `await` or the callback `.then()` mechanism can be implemented using Go channels.

# Lazy Evaluation

Lazy Evaluation is a technique to reduce memory usage. If you ever use one of Python's generator functions like `range(n)` or the keyword `yield`, that is a kind of Lazy Evaluation. 

For example, a function that returns a list of objects can be implemented as a function that returns a channel. The caller can then receive from the channel to get the objects one by one. This way, the caller doesn't need to wait for the whole list to be generated before it can start processing the objects.

# Benefits

## Asynchronous Communication: 

By returning a channel, the function allows the caller to communicate asynchronously with the underlying goroutines. This enables concurrent processing of data without blocking or waiting for each operation to complete.

## Decoupling: 

The pattern provides a decoupled way of interacting with concurrent components. The caller does not need to have knowledge of the internal workings of the goroutines. It only needs to interact with the returned channel, allowing for loose coupling and easier maintenance of the codebase.

## Flexibility: 

The pattern allows for flexibility in handling concurrent operations. The caller can choose when and how to process the received values, enabling custom logic for parallel execution, throttling, or synchronization.

## Concurrency Control: 

By utilizing channels, the pattern provides a built-in synchronization mechanism. The caller can use channels to coordinate goroutines, wait for results, or manage the overall flow of the concurrent operations.
