package main

import "fmt"
import "runtime"

const goal = 200

func countOptions(total int, options []int) int {

	// We got a solution
	if total == goal {
		return 1
	}

	// We went to far
	if total > goal {
		return 0
	}

	// We didn't go far enough
	if len(options) == 0 {
		return 0
	}

	// Pick the first option and continue to allow that option as a choice
	// Then skip the first option and recurse
	return countOptions(total+options[0], options) + countOptions(total, options[1:])
}

func main() {
	runtime.GOMAXPROCS(1)

	options := []int{200, 100, 50, 20, 10, 5, 2, 1}
	fmt.Printf("%v", countOptions(0, options))
}
