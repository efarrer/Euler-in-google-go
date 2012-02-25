package main

import "fmt"
import "runtime"
import "strconv"

func isPalinDromic(txt string) bool {
	s := 0
	e := len(txt) - 1

	for s < e {
		if txt[s] != txt[e] {
			return false
		}
		s++
		e--
	}

	return true
}

func main() {
	runtime.GOMAXPROCS(1)

	const max = 1000000

	count := int64(0)
	for i := int64(0); i != max; i++ {
		if isPalinDromic(strconv.FormatInt(i, 10)) {
			if isPalinDromic(strconv.FormatInt(i, 2)) {
				print(i, "\n")
				count += i
			}
		}
	}

	fmt.Printf("%v", count)
}
