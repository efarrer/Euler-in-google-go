package main

import "fmt"
import "runtime"
import "strconv"
import "math"
import "regexp"

func toDigits(num int64) []int64 {
	str := strconv.FormatInt(num, 10)
	digits := make([]int64, len(str))
	for i := 0; i != len(str); i++ {
		digits[i], _ = strconv.ParseInt(str[i:i+1], 10, 0)
	}
	return digits
}

func main() {
	runtime.GOMAXPROCS(1)

	sum := int64(0)
	for n := int64(10); ; n++ {
		total := int64(0)
		digits := toDigits(n)
		for i := 0; i < len(digits); i++ {
			total += int64(math.Pow(float64(digits[i]), 5))
		}
		if total == n {
			sum += total
		}

		// See if we have all 9's
		nStr := strconv.FormatInt(n, 10)
		totalStr := strconv.FormatInt(total, 10)

		// If the biggest possible number for a digit range has fewer digits
		// than the total then we are done
		if len(totalStr) < len(nStr) {
			if matched, _ := regexp.Match("^9+$", []byte(nStr)); matched {
				break
			}
		}
	}
	fmt.Printf("%v", sum)
}
