package main

import "fmt"
import "runtime"
import "strconv"
import "regexp"

func charsToInt64(str string) []int64 {
	ints := make([]int64, len(str))
	for i := 0; i < len(str); i++ {
		ints[i], _ = strconv.ParseInt(str[i:i+1], 10, 64)
	}

	return ints
}

var cache []int64

func init() {
	cache2 := make([]int64, 10)
	for i := int64(0); i != 10; i++ {
		cache2[i] = fact(i)
	}
	cache = cache2
}

func fact(d int64) int64 {
	if nil != cache {
		return cache[d]
	}
	if 0 == d {
		return 1
	}
	return d * fact(d-1)
}

func main() {
	runtime.GOMAXPROCS(1)

	total := int64(0)
	for i := int64(3); ; i++ {
		strI := strconv.FormatInt(i, 10)
		digits := charsToInt64(strI)
		sum := int64(0)
		for _, d := range digits {
			sum += fact(d)
		}
		if sum == i {
			total += sum
		}

		// If the number is all 9's and it's bigger than the sum(fact(digits))
		// then we won't ever find another match
		if matched, _ := regexp.Match("^9+$", []byte(strI)); matched {
			if i > sum {
				break
			}
		}
	}

	fmt.Printf("%v", total)
}
