package main

import "fmt"
import "runtime"
import "math"

func main() {
	runtime.GOMAXPROCS(1)

	results := make(map[int]int)
	maxKey := 0
	maxValue := 0

	// The a side must be less than half of c
	for a := 1; a < 500; a++ {
		print(a, "\n")
		// The b side must be less than half of c
		for b := 1; b < 500; b++ {
			c2 := a*a + b*b
			c := int(math.Sqrt(float64(c2)))

			// B is too big
			if c > 1000 {
				break
			}
			// Doesn't have integral sides
			if c*c != c2 {
				continue
			}

			perimiter := a + b + c
			results[perimiter]++
			if results[perimiter] > maxValue {
				maxValue = results[perimiter]
				maxKey = perimiter
			}
		}
	}
	print("done\n")
	fmt.Printf("%v", maxKey)
}
