package main

import "fmt"
import "runtime"
import "strconv"

func remove(st string, i int) (string, string) {
	rest := ""
	removed := st[i : i+1]
	if i != 0 {
		rest += st[0:i]
	}
	if i != len(st) {
		rest += st[i+1:]
	}

	return removed, rest
}

func permutations(st string) []string {

	res := []string{}

	if st == "" {
		return []string{""}
	}

	for i := 0; i != len(st); i++ {
		removed, rest := remove(st, i)
		restPerm := permutations(rest)
		for j := 0; j != len(restPerm); j++ {
			res = append(res, removed+restPerm[j])
		}
	}

	return res
}

func toInt(num string) int64 {
	i, e := strconv.ParseInt(num, 10, 64)
	if nil != e {
		panic(e)
	}
	return i
}

func main() {
	runtime.GOMAXPROCS(1)

	pandigital := "9876543210"
	allPandigitals := permutations(pandigital)

	var divisors = map[int]int64{1: 2, 2: 3, 3: 5, 4: 7, 5: 11, 6: 13, 7: 17}

	sum := int64(0)
	for _, numStr := range allPandigitals {
		passes := true
		for i := 1; i < 8; i++ {
			if toInt(numStr[i:(i+3)])%divisors[i] != 0 {
				passes = false
				break
			}
		}
		if passes {
			sum += toInt(numStr)
		}
	}
	fmt.Printf("%v", sum)
}
