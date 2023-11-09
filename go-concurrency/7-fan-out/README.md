# Fan out

## Problem

You have a channel, and you want to send the messages to multiple goroutines. How do you do that?

## Solution

Use a select statement to receive from the channel, and send to multiple channels.