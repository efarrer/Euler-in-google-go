package main

import (
	"fmt"
	"math/big"
	"runtime"
)

const digits = 1000

func fib(ith int, prev, prevprev *big.Int) int {
	fibval := big.NewInt(0)
	fibval.Add(prev, prevprev)
	if digits == len(fibval.String()) {
		return ith
	}
	return fib(ith+1, fibval, prev)
}

func findFirstFibWithXDigits() int {
	return fib(2, big.NewInt(1), big.NewInt(0))
}

func main() {
	runtime.GOMAXPROCS(1)

	fmt.Printf("%v", findFirstFibWithXDigits())
}
