package main

import "fmt"

func main() {
	var backtwo int64 = 0
	var backone int64 = 1
	var fib int64 = 0
	var total int64 = backtwo // backtwo is even backone is not
	const max = 4000000

	for {
		fib = backtwo + backone
		backtwo = backone
		backone = fib

		if fib >= max {
			break
		}

		if 0 == (fib % 2) {
			total += fib
		}
	}

	fmt.Printf("%v", total)
}
