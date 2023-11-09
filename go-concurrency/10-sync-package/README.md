# `sync` package 

This package contains multiple synchronization mechanisms.

## Mutex

`Mutex` is a mutual exclusion lock. It is used to synchronize access to a shared resource. Only one goroutine can hold the lock at a time. Other goroutines will wait until the lock is released.

## RWMutex

`RWMutex` is a reader/writer mutual exclusion lock. It is used to synchronize access to a shared resource. Multiple readers can hold the lock at the same time. Only one writer can hold the lock at a time. Other goroutines will wait until the lock is released.

## WaitGroup

`WaitGroup` is used to wait for a collection of goroutines to finish. The main goroutine calls `Add()` to set the number of goroutines to wait for. Then each of the goroutines runs and calls `Done()` when finished. At the same time, `Wait()` can be used to block until all goroutines have finished.

## Once

`Once` is used to perform initialization exactly once. It is used to lazily initialize a resource. The `Do()` method takes a function as an argument. The function is executed only once, regardless of how many times `Do()` is called.

## Cond

`Cond` is a condition variable. It is used to block or wake goroutines. The `Wait()` method blocks the goroutine until `Signal()` or `Broadcast()` is called. The `Signal()` method wakes one goroutine waiting on the condition variable. The `Broadcast()` method wakes all goroutines waiting on the condition variable.

## Pool

`Pool` is used to manage a pool of resources. It is used to reuse resources instead of creating and destroying them. The `Get()` method returns a resource from the pool. The `Put()` method returns a resource to the pool.

## Map 

`Map` is a concurrent map. It is used to store key-value pairs. The `Load()` method returns the value for a key. The `Store()` method stores a value for a key. The `Delete()` method deletes a key-value pair. The `Range()` method iterates over all key-value pairs.

# `sync/atomic` package

`Atomic` is used to perform atomic operations on memory. It is used to synchronize access to a shared resource. The `AddInt64()` method atomically adds delta to *addr and returns the new value. The `LoadInt64()` method atomically loads *addr. The `StoreInt64()` method atomically stores val into *addr.

# When to use `sync` package

Don't use `sync` package unless you really have to. Use `channel` and `select` instead.

