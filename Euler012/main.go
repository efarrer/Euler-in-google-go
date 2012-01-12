package main

import "fmt"
import "runtime"

const max = 500

type divisors struct {
	number int64
	count  int
}

func divisorCount(item int64) int {
	count := 0

	for i := int64(1); i <= item; i++ {
		if 0 == (item % i) {
			count++
		}
	}

	return count
}

func countDivisors(recvChan chan int64, sendChan chan *divisors) {
	for {
		count := 0
		number := <-recvChan
		for i := int64(1); i <= number; i++ {
			if 0 == (number % i) {
				count++
			}
		}

		sendChan <- &divisors{number, count}
	}
}
func main() {
	runtime.GOMAXPROCS(8)

	goRoutines := runtime.GOMAXPROCS(-1)

	sendChan := make(chan int64)
	recvChan := make(chan *divisors)

	// Create the divisors threads
	for i := 0; i < goRoutines; i++ {
		go countDivisors(sendChan, recvChan)
	}

	outstandingOperations := 0
	var currentDivisor *divisors = nil

	i := int64(1)
	j := i
the_loop:
	for {
		select {
		case sendChan <- j:
			outstandingOperations++
			i++
			j += i
		case currentDivisor = <-recvChan:
			outstandingOperations--
			if currentDivisor.count > max {
				break the_loop
			}
		}
	}

	// At this point we've found one triangle number with at least 500 divisors,
	// but we may not have found the first one since we're processing many in
	// parallel. So get the rest and pick the best
	for outstandingOperations > 0 {
		otherDivisor := <-recvChan
		outstandingOperations--

		// It meets the criteria
		if otherDivisor.count > max {
			// It smaller than the first found one
			if otherDivisor.number < currentDivisor.number {
				currentDivisor = otherDivisor
			}
		}
	}
	fmt.Printf("%v", currentDivisor.number)
}
