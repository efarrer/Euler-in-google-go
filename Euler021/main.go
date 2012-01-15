package main

import "fmt"
import "runtime"

func sum(nums []int) int {
	total := 0
	for i := 0; i != len(nums); i++ {
		total += nums[i]
	}
	return total
}

func properDivisors(num int) []int {
	divisors := make([]int, 0, 10)
	for i := 1; i <= (num / 2); i++ {
		if 0 == (num % i) {
			divisors = append(divisors, i)
		}
	}

	return divisors
}

func getAmicablePair(num int) (int, bool) {
	pair := sum(properDivisors(num))

	// Amicable pair can't be itself
	if pair == num {
		return 0, false
	}

	orig := sum(properDivisors(pair))

	// If it doesn't transform back to us then it's not amicable
	if orig != num {
		return 0, false
	}

	// Yea we're amicable
	return pair, true
}

func main() {
	runtime.GOMAXPROCS(1)

	amicablePairs := map[int]bool{}

	for i := 1; i < 10000; i++ {
		pair, amm := getAmicablePair(i)
		if amm {
			amicablePairs[i] = true
			amicablePairs[pair] = true
		}
	}

	total := 0
	for key, _ := range amicablePairs {
		total += key
	}

	fmt.Printf("%v", total)
}
