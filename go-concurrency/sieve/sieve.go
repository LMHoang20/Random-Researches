package main

import (
	"fmt"
	"sync"
	"time"
)

func isPrime(MaxN int) []int {
	isPrime := make([]int, MaxN)

	for i := 2; i < MaxN; i++ {
		isPrime[i] = 1
	}

	return isPrime
}

func concurrentSieve(MaxN int) {
	isPrime := isPrime(MaxN)

	waitGroup := sync.WaitGroup{}
	for i := 2; i < MaxN; i++ {
		waitGroup.Add(1)
		go sieve(i, isPrime, MaxN, &waitGroup)
	}
	waitGroup.Wait()
}

func nonConcurrentSieve(MaxN int) {
	isPrime := isPrime(MaxN)

	for i := 2; i < MaxN; i++ {
		sieve(i, isPrime, MaxN, nil)
	}
}

func sieve(i int, isPrime []int, MaxN int, waitGroup *sync.WaitGroup) {
	if waitGroup != nil {
		defer waitGroup.Done()
	}
	for j := 2 * i; j < MaxN; j += i {
		isPrime[j] = 0
	}
}

func timeIt(f func()) {
	start := time.Now()
	f()
	fmt.Println(time.Since(start))
}

func main() {
	testMaxN := []int{100, 1000, 10000, 100000, 1000000}

	for _, MaxN := range testMaxN {
		fmt.Printf("MaxN = %d\n", MaxN)
		fmt.Println("Concurrent:")
		timeIt(func() {
			concurrentSieve(MaxN)
		})
		fmt.Println("Non-Concurrent:")
		timeIt(func() {
			nonConcurrentSieve(MaxN)
		})
		fmt.Println()
	}
}
