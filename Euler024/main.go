package main

import "fmt"
import "runtime"
import "sort"
import "strconv"

func charsToInt(str string) []int {
	ints := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		ints[i], _ = strconv.Atoi(str[i : i+1])
	}

	return ints
}

func intsToString(ints []int) string {
	str := ""

	for _, i := range ints {
		str = str + strconv.Itoa(i)
	}
	return str
}

func max(lst []int) int {
	res := -1
	for _, i := range lst {
		if i > res {
			res = i
		}
	}
	return res
}

func min(lst []int) int {
	res := 10000
	for _, i := range lst {
		if i < res {
			res = i
		}
	}
	return res
}

func filter(lst []int, fn func(int) bool) []int {
	res := []int{}

	for _, val := range lst {
		if fn(val) {
			res = append(res, val)
		}
	}

	return res
}

func remove(lst []int, item int) []int {
	res := []int{}
	for _, val := range lst {
		if val != item {
			res = append(res, val)
		}
	}

	return res
}

func lexicographicPermutation(str string) string {
	digits := charsToInt(str)
	maxIndex := len(digits) - 1
	prevValue := digits[maxIndex]
	currValue := prevValue

	for i := maxIndex; i != -1; i-- {
		currValue = digits[i]
		if i == 0 && currValue == max(digits) {
			return str
		}
		if currValue < prevValue {
			prefix := digits[:i]
			suffix := digits[i:]
			nextStart := min(filter(suffix, func(x int) bool { return x > currValue }))
			suffix = remove(suffix, nextStart)
			sort.IntSlice(suffix).Sort()
			return intsToString(prefix) + strconv.Itoa(nextStart) + intsToString(suffix)
		} else {
			prevValue = currValue
		}
	}

	return str
}

func main() {
	runtime.GOMAXPROCS(1)
	start := "0123456789"
	for i := 0; i < 999999; i++ {
		start = lexicographicPermutation(start)
	}
	fmt.Printf("%v", start)
}
