package main

import "fmt"

const MAX int64 = 20

func testAnswer(value int64) bool {
	for j := (MAX - 1); j >= int64(1); j-- {
		if 0 != (value % j) {
			return false
		}
	}
	return true
}

func main() {
	for i := MAX; ; i += MAX {

		if testAnswer(i) {
			fmt.Printf("%v", i)
			return
		}
	}
}
