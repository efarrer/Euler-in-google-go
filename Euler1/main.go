package main

import "fmt"

func main() {
	const max = 1000
	total := 0

	for i := 0; i < max; i++ {
		if 0 == (i%3) || 0 == (i%5) {
			total += i
		}
	}
	fmt.Printf("%v", total)
}
