# Fan In Pattern

## Problem

You have multiple goroutines, each sending messages to its own channel. But some where in your project, you want to merge them into one channel. How do you do that?

## Solution

Create a function that takes multiple channels as arguments, and returns a channel. Inside the function, create a goroutine for each channel, and send the messages to the returned channel.

![](https://go.dev/talks/2012/concurrency/images/gophermegaphones.jpg)

