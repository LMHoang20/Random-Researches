package main

import (
	"fmt"
	"time"
)

func Generator(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func chanSieve(n int) {
	ch := make(chan int)
	go Generator(ch)
	for {
		prime := <-ch
		if prime >= n {
			break
		}
		// fmt.Print(prime, " ")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

func nonConcurrentSieve(n int) {
	primes := []int{}

	for i := 2; i < n; i++ {
		isPrime := true
		for _, prime := range primes {
			if i%prime == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			// fmt.Print(i, " ")
			primes = append(primes, i)
		}
	}
}

func timeit(f func()) {
	start := time.Now()
	f()
	fmt.Println(time.Since(start))
}

func main() {
	testMaxN := []int{100, 1000, 10000, 100000}

	for _, MaxN := range testMaxN {
		fmt.Printf("MaxN = %d\n", MaxN)
		fmt.Print("Concurrent: ")
		timeit(func() {
			chanSieve(MaxN)
		})
		fmt.Print("Non-Concurrent: ")
		timeit(func() {
			nonConcurrentSieve(MaxN)
		})
		fmt.Println()
	}
}
