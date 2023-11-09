# Goroutine

A **Goroutine** is a lightweight and concurrent execution unit in Golang, allowing tasks to be executed concurrently within a single Go program.

A **Goroutine** in Golang can be created by the use of the keyword `go`. The syntax to create a **Goroutine** is as follows:

```go
go function_name()
```

# Communicating Sequential Processes

1. Each process is built for *sequential* exexcution.

2. Data is *communicated* between processes via *channels*. No shared state.

3. Scale by adding more processes.

# Internal Working

![Goroutine](https://www.ardanlabs.com/images/goinggo/94_figure2.png)

- **G**: Goroutine
- **M**: Thread (Machine)
- **P**: Logical Processor
- **GRQ**: Global Runtime Queue
- **LRQ**: Local Runtime Queue

# Goroutine vs Thread

A **Goroutine** is basically a light-weight thread implemented in the application level.

|| Thread | Goroutine |
|---|--------|-----------|
| Management | OS | Go Runtime |
| Scheduler | Preemptive | Cooperating |
| Stack size | MB of memory, depends on hardware, static | 2KB of memory, allocated dynamically |
| Creation / Context switch | Expensive Syscall to Kernel mode | Everything in User mode |
| Optimization | Kernel-level, have to be robust | Application-level, can do more optimization |
| Number | 1000s | 100000s |


# State

### Waiting: 

This means the Goroutine is stopped and waiting for something in order to continue. This could be for reasons like waiting for the operating system (system calls) or synchronization calls (atomic and mutex operations). These types of latencies are a root cause for bad performance.

### Runnable: 

This means the Goroutine wants time on an M so it can execute its assigned instructions. If you have a lot of Goroutines that want time, then Goroutines have to wait longer to get time. Also, the individual amount of time any given Goroutine gets is shortened as more Goroutines compete for time. This type of scheduling latency can also be a cause of bad performance.

### Executing: 

This means the Goroutine has been placed on an M and is executing its instructions. The work related to the application is getting done. This is what everyone wants.

# Scheduler

## Context-switch

**Goroutines** are scheduled by a *cooperating* scheduler managed by Go runtime. This scheduler runs in **user space**, above the kernel. Being a *preemptive* scheduler, OS threads can stop and switch at any point during execution. *cooperating* scheduling means that Goroutines can only switch during safe points.

There are four classes of events that occur in Go programs that allow the scheduler to make scheduling decisions. This doesn’t mean it will always happen on one of these events. It means the scheduler gets the opportunity.

### Keyword go

Whenever we start a goroutine, the scheduler can use    this opportunity to switch.

### Garbage collection

Go Garbage collection (GC) runs using its own set of Goroutines, and those Goroutines need its thread to run. So the GC creates a lot of chaos when it runs. During this time, the scheduler can switch off the Goroutines that uses the heap and switch on the Goroutines that don’t.

### System calls

If a system call is asynchronous (M is not blocked), something called the network poller can be used to process the system call more efficiently. This is accomplished by using `kqueue (MacOS)`, `epoll (Linux)` or `iocp (Windows)` within these respective OS’s. The scheduler will switch that goroutine off to wait for the network poller and switch on another Goroutine to run on the same thread.

If a system call is synchoronous (M is blocked), the scheduler move the rest of the goroutines in the thread's LRQ to a new thread. After the system call is done, the caller goroutine will also be moved to that new thread. The original thread will be kept for a while and (maybe) reused if similar system calls are made. This is called **thread caching**.

### Blocking operations

If an atomic, mutex, or channel operation call will cause the Goroutine to block, the scheduler can context-switch a new Goroutine to run. Once the Goroutine can run again, it can be re-queued and eventually context-switched back.

## Work stealing

The last thing you want is for a thread to be idle. Because then, the OS will context-switch the thread off the CPU. Even though there are still goroutines that are runnable.

So, when a thread is idle, it will look to execute Goroutines waiting in GRQ. If there are nothing in GRQ, it will look to steal work from other threads' LRQ. This is called **work stealing**. 
