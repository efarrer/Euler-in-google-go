package main

import "fmt"
import "runtime"
import "strconv"
import "strings"

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

func toInt(num string) int64 {
	i, e := strconv.ParseInt(num, 10, 64)
	if nil != e {
		panic(e)
	}
	return i
}

func main() {
	const cpus = 8
	runtime.GOMAXPROCS(cpus)

	taskCh := make(chan int64)
	resultCh := make(chan int64)

	for c := 0; c != cpus; c++ {
		go func() {
			for {
				num := <-taskCh
				str := strconv.FormatInt(num, 10)
				for mult := int64(2); true; mult++ {
					str += strconv.FormatInt(num*mult, 10)
					length := len(str)
					if length == 9 {
						if isPandigitalOneToNine(str) {
							resultCh <- toInt(str)
						} else {
							resultCh <- 0
						}
						break
					} else if length > 9 {
						resultCh <- 0
						break
					}
				}
			}
		}()
	}

	// The minimal result is biggestInput . (2*biggestInput) and it must have
	// less digits that 987654321. So pick a limit that will just barely exceed
	// 9 digits (50000 . 100000) is 11 digits which is close enough
	const biggestInput = 50000

	biggest := int64(0)
	outstanding := 0
	num := int64(1)
	for num <= biggestInput { // Push requests and handle results
		select {
		case taskCh <- num:
			outstanding++
			num++
		case result := <-resultCh:
			outstanding--
			if result > biggest {
				biggest = result
				print("Biggest: ", biggest, "\n")
			}
		}
	}
	for outstanding > 0 { // Handle outstanding results
		result := <-resultCh
		outstanding--
		if result > biggest {
			biggest = result
			print("Biggest: ", biggest, "\n")
		}
	}
	fmt.Printf("%v", biggest)
}
