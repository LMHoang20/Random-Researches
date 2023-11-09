# Drawbacks of Goroutines, Channels, and the CSP model

Goroutines and Channels can get you a long way in Go concurrency, but sometimes they are not enough. 

## Goroutines and Channels

`channel` in Golang is very powerful, but it does have some limitations:

- `channel` pass around copies of objects... not the best, performance-wise.

- Some structure are naturally shared like `caches` or `registries`

- Even though most of `sync` package can be implemented using `channel`, it is not as readable and easy to understand.

## Other potential drawbacks

### Lack of control over low-level details. 

Go's concurrency model abstracts away many low-level details, which can be a double-edged sword. While this abstraction makes concurrent programming more straightforward, it also removes some control from the developers. For instance, the Go runtime decides how to schedule goroutines, and developers have little control over this process.

### Potential for goroutine leaks. 

Goroutines are cheap to create, which can lead to their overuse. If a program starts a goroutine and receives nothing from the channel it sends on, it leaks. Over time, these leaks can consume all available memory, leading to program crashes.

### Difficulty handling errors. 

Error handling in concurrent Go programs can be tricky. If an error occurs in a goroutine, it can't be returned to the caller because it has a different call stack. It requires developers to use other mechanisms, like channels or shared memory, to propagate errors.

### Error propagation for nested goroutines. 

When goroutines spawn other goroutines, managing their lifetimes and propagating errors can become complex. The parent goroutine must keep track of its children and ensure they handle errors correctly.