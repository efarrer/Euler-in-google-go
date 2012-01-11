package main

import "fmt"
import "runtime"

const MAX int64 = 2000000

func createNewPrimeFilter(prime int64, ch, res chan int64) {

	// The next expected multiple of 'prime'
	expected := prime

	var nextCh chan int64 = nil

	for {
		next := <-ch

		// Jump to the next expected prime
		for expected < next {
			expected += prime
		}

		// If we didn't get a multiple of our prime then it's a new prime
		if next != expected {
			// Pass it on to the next filter
			if nil == nextCh {
				nextCh = make(chan int64)
				res <- next
				go createNewPrimeFilter(next, nextCh, res)
			} else {
				nextCh <- next
			}
		}
	}
}

func main() {
	runtime.GOMAXPROCS(4)

	const start = int64(2)

	ch := make(chan int64)
	res := make(chan int64)

	go createNewPrimeFilter(start, ch, res)

	// Pump to the filters
	go func() {
		for i := start; ; i++ {
			ch <- i
		}
	}()

	// Sum the primes
	total := start
	for {
		cur := <-res

		if cur > MAX {
			fmt.Printf("%v\n", total)
			return
		}
		total += cur
	}

}
