package main

import "fmt"
import "runtime"
import "math/big"

type results struct {
	a int
	b int
	n int
}

func main() {
	const procs = 8
	runtime.GOMAXPROCS(procs)

	const min = -999
	const max = 1000

	// Goroutine for aggregating the results
	bestCh := make(chan *results)

	// Goroutine for calculating the results
	for a := min; a < max; a++ {
		go func(aa int) {
			for b := min; b < max; b++ {
				n := 0
				for {
					if !big.NewInt(int64((n * n) + (aa * n) + b)).ProbablyPrime(1000) {
						bestCh <- &results{aa, b, n - 1}
						break
					} else {
						n++
					}
				}
			}
		}(a)
	}

	best := &results{0, 0, 0}
	for a := min; a < max; a++ {
		for b := min; b < max; b++ {
			this := <-bestCh
			if this.n > best.n {
				best = this
			}
		}
	}
	fmt.Printf("%v", best.a*best.b)
}
