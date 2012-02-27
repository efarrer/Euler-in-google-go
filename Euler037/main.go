package main

import "fmt"
import "runtime"
import "math/big"
import "strconv"

func isPrime(num int64) bool {
	return big.NewInt(num).ProbablyPrime(100000)
}

// We know this is safe
func toInt(num string) int64 {
	i, e := strconv.ParseInt(num, 10, 64)
	if nil != e {
		panic(e)
	}
	return i
}

func isTruncatiblePrime(num string) bool {
	numInt := toInt(num)
	// Base case is single digit number
	if len(num) == 1 {
		return isPrime(numInt)
	}

	// Check the number itself
	if !isPrime(numInt) {
		return false
	}

	head := num[:len(num)-1]
	for len(head) > 0 {
		headInt := toInt(head)
		if !isPrime(headInt) {
			return false
		}
		head = head[:len(head)-1]
	}

	tail := num[1:]
	for len(tail) > 0 {
		tailInt := toInt(tail)
		if !isPrime(tailInt) {
			return false
		}
		tail = tail[1:]
	}

	return true
}

func main() {
	runtime.GOMAXPROCS(1)

	// The list of digits that can be used
	// These are the only single digit primes
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	count := int64(0)
	sum := int64(0)

	// The numbers to check
	workQueue := []string{"2", "3", "5", "7"}

	for count < 11 {
		// Pop an item off of the work queue
		item := workQueue[0]
		workQueue = workQueue[1:]

		for _, digit := range digits {
			totest := item + digit
			print("\tTest: ", totest, "\n")
			if isTruncatiblePrime(totest) {
				sum += toInt(totest)
				count++
				print("Found: ", totest, "\n")
			}
			workQueue = append(workQueue, totest)
		}
	}
	fmt.Printf("%v", sum)
}
