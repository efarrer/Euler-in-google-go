package main

import "fmt"
import "runtime"

const abundant_ceiling = 28124

func properDivisors(num int) []int {
	divisors := make([]int, 0, 10)
	for i := 1; i <= (num / 2); i++ {
		if 0 == (num % i) {
			divisors = append(divisors, i)
		}
	}

	return divisors
}

func sum(nums []int) int {
	total := 0
	for i := 0; i != len(nums); i++ {
		total += nums[i]
	}
	return total
}

type abundantResponse struct {
	num      int
	abundant bool
}

func isAbundant(num int, ch chan *abundantResponse) {
	divisors := properDivisors(num)
	abundant := sum(divisors) > num
	ch <- &abundantResponse{num, abundant}
}

const done = -1

func calculateSums(index int, abundantNumbers []int, sumCh chan int) {
	for i := index; i < len(abundantNumbers); i++ {
		sumCh <- abundantNumbers[index] + abundantNumbers[i]
	}
	sumCh <- done
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	abundantNumbers := make([]int, 0, abundant_ceiling)

	// Calculate all abundant numbers
	abundantCh := make(chan *abundantResponse)
	for i := 1; i < abundant_ceiling; i++ {
		go isAbundant(i, abundantCh)
	}
	for i := 1; i < abundant_ceiling; i++ {
		response := <-abundantCh
		if response.abundant {
			abundantNumbers = append(abundantNumbers, response.num)
		}
	}

	// Spawn one goroutine for each abundant number to calculate the sums
	sumCh := make(chan int)
	for i, _ := range abundantNumbers {
		go calculateSums(i, abundantNumbers, sumCh)
	}

	outStanding := len(abundantNumbers)
	// Will store all numbers that aren't the sum of an abundant number
	sumOfAbundant := make([]bool, abundant_ceiling-1)
	for outStanding > 0 {
		res := <-sumCh
		if done == res {
			outStanding--
		} else if res < len(sumOfAbundant) {
			sumOfAbundant[res] = true
		}
	}

	total := 0
	for i, isSum := range sumOfAbundant {
		if !isSum {
			total += i
		}
	}

	fmt.Printf("%v", total)
}
