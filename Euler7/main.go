package main

import "fmt"
import "os"
import "runtime"

const ithPrime = int64(10001)

func createNewPrimeFilter(prime, ith int64, ch chan int64) {

	if ithPrime == ith {
		fmt.Printf("%v", prime)
		os.Exit(0)
	}

	// The next expected multiple of 'prime'
	expected := prime

	var nextCh chan int64 = nil

	for {
		next := <-ch

		// Jump to the next expected prime
		for expected < next {
			expected += prime
		}
		// If we got a multiple of our prime swallow it
		if next != expected {
			// Pass it on to the next filter
			if nil == nextCh {
				nextCh = make(chan int64)
				go createNewPrimeFilter(next, ith+1, nextCh)
			}
			nextCh <- next
		}
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	ch := make(chan int64)
	go createNewPrimeFilter(2, 1, ch)

	for i := int64(2); ; i++ {
		ch <- i
	}
}
