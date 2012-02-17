package main

import "fmt"
import "runtime"
import "strings"
import "strconv"

func isPandigitalOneToNine(number string) bool {
	if len(number) != 9 {
		return false
	}

	for i := int64(1); i <= 9; i++ {
		if !strings.Contains(number, strconv.FormatInt(i, 10)) {
			return false
		}
	}
	return true
}

func main() {
	runtime.GOMAXPROCS(1)

	tosum := map[int64]bool{}

	for a := int64(1); ; a++ {
		// For the smallest b (b == 1) case the result will be the same length
		// as digits(a) so we can stop when a > 9999 (4+1+4 == 9 so it's OK but 5+1+5
		// is too far)
		if a > 9999 {
			break
		}
		for b := int64(a); ; b++ {
			// Same as above
			if b > 9999 {
				break
			}

			c := a * b
			aStr := strconv.FormatInt(a, 10)
			bStr := strconv.FormatInt(b, 10)
			cStr := strconv.FormatInt(c, 10)
			if isPandigitalOneToNine(aStr + bStr + cStr) {
				tosum[c] = true
			}
		}
	}

	sum := int64(0)
	for key, _ := range tosum {
		sum += key
	}

	fmt.Printf("%v", sum)
}
