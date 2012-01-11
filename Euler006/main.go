package main

import "fmt"

const MAX int64 = 100

func main() {

	sumOfSquares := int64(0)
	for i := int64(1); i <= MAX; i++ {
		sumOfSquares += (i * i)
	}

	squareOfSums := int64(0)
	for i := int64(1); i <= MAX; i++ {
		squareOfSums += i
	}
	squareOfSums *= squareOfSums

	fmt.Printf("%v", squareOfSums-sumOfSquares)
}
