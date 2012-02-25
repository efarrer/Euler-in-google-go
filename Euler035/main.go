package main

import "fmt"
import "runtime"
import "strconv"
import "math/big"

func rotations(str string) []string {
	strs := []string{}

	if "" == str {
		return []string{""}
	}

	for i := 0; i != len(str); i++ {
		second := str[0:i]
		first := str[i:]
		strs = append(strs, first+second)
	}

	return strs
}

func rotationsInt(num int64) []int64 {
	strs := rotations(strconv.FormatInt(num, 10))
	nums := []int64{}
	for _, str := range strs {
		n, _ := strconv.ParseInt(str, 10, 64)
		nums = append(nums, n)
	}
	return nums
}

func isPrime(num int64) bool {
	return big.NewInt(num).ProbablyPrime(100000)
}

func main() {
	runtime.GOMAXPROCS(8)

	const max = 1000000

	// Calculate all the primes
	ch := make(chan int64)
	for i := int64(2); i != max; i++ {
		go func(j int64) {
			allPrimes := false
			if isPrime(j) {
				allPrimes = true
				rots := rotationsInt(j)
				for _, r := range rots {
					if !isPrime(r) {
						allPrimes = false
						break
					}
				}
			}
			if allPrimes {
				ch <- j
			} else {
				ch <- 0
			}
		}(i)
	}
	count := 0
	for i := int64(2); i != max; i++ {
		p := <-ch
		if p != 0 {
			print(p, "\n")
			count++
		}
	}
	fmt.Printf("%v", count)
}
