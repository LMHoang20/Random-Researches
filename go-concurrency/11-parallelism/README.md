# Concurrency vs Parallelism

- Concurrency is quickly switching between tasks. Doing multiple things but once at a time.

- Parallelism is doing multiple things at the same time.

# How to achieve Parallelism

# GOMAXPROCS

The `GOMAXPROCS` environment variable allows you to limit the number of operating system threads that can execute user-level Go code simultaneously. Starting with Go version 1.5, the default value of `GOMAXPROCS` should be the number of logical cores available in your UNIX machine.

If you decide to assign a value to `GOMAXPROCS` that is less than the number of the cores in your UNIX machine, you might affect the performance of your program. However, using a `GOMAXPROCS` value that is larger than the number of the available cores will not necessarily make your Go programs run faster.

You can programmatically discover the value of the `GOMAXPROCS` environment variable; the relevant code can be found ...