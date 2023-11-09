# Quit channel

## Problem

Normally, you don't want to close the channel from the receiver side because the sender side will panic. But sometimes, you want to close the channel from the receiver side. How do you do that?

## Solution

Create a channel that is used to signal the sender to close the channel. The sender can then close the channel when it receives the signal.