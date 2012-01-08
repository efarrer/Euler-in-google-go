package main

import "fmt"
import "math"
import "math/big"

func main() {
	const number int64 = 600851475143
	var maxFactor int64 = int64(math.Sqrt(float64(number)))

	for factor := maxFactor; factor >= 2; factor-- {
		// Make sure it divides
		if 0 == (number % factor) {
			// Make sure it's prime
			if big.ProbablyPrime(big.NewInt(factor), 10000) {
				fmt.Printf("%v", factor)
				return
			}
		}
	}
}
