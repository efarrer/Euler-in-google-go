package main

import "fmt"
import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	var start int = 0
	var end int = len(str) - 1

	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func main() {
	const maxVal = 999
	const minVal = 99
	var x int = maxVal
	var y int = minVal

	largest := 0

	for x = maxVal; x > minVal; x-- {
		for y = maxVal; y > minVal; y-- {
			cur := x * y

			// Since y is going down at this point we won't find a bigger one so
			// we can safely just quit this loop and try the next smallest x
			if cur < largest {
				break
			}

			if isPalindrome(cur) {
				largest = cur
			}
		}
	}
	fmt.Printf("%v", largest)
}
