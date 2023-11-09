# Select

`select` is a keyword in **Golang**. It is used to choose between multiple send/receive channel operations. 

The `select` statement blocks until one of the send/receive operation is ready. 

If multiple operations are ready, one of them is chosen at random. If nothing is ready, the `default` case is executed.

```go

func combineSearchResult(...) {
    for {
        select {
        case data := <-textSearchService:
            addDataToResult(data)
        case data := <-imageSearchService:
            addDataToResult(data)
        case <-time.After(20 * time.Second):
            healthCheck()
        }
    }
}


```

The syntax is similar to `switch`, but instead of comparing values, it compares channels. The `default` case is optional.

`select` is often used with `time.After()` to implement a timeout mechanism.

## Fan In pattern

`select` is the better way to implement the fan-in pattern. It is more readable and easier to understand.
