package main

import "fmt"
import "runtime"

func main() {
	const MAX = 1000

	runtime.GOMAXPROCS(1)

	for c := MAX; c >= 0; c-- {
		for b := c - 1; b >= 0; b-- {
			if b+c > MAX {
				continue
			}
			for a := b - 1; a >= 0; a-- {
				if a+b+c != MAX {
					continue
				}

				if (a*a)+(b*b) == (c * c) {
					fmt.Printf("%v", a*b*c)
					return
				}
			}
		}
	}
}
