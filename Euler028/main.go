package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(1)

	const size = 1001 * 1001

	// The grand total of the diagonals
	total := 0
	// How many to skip before we hit a diagonal
	skipCount := 0
	// How many we've skipped so far
	skipped := 0
	// How may of the diagonals for a particular radius we've hit
	diagsHit := 3 // We start with 3 because the center square really covers all four diagonals
	for count := 1; count <= size; count++ {
		if skipped == skipCount {
			skipped = 0
			total += count
			diagsHit++
			if 4 == diagsHit {
				skipCount += 2
				diagsHit = 0
			}
		}
		skipped++
	}
	fmt.Printf("%v", total)
}
