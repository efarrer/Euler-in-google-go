package main

import "fmt"
import "runtime"
import "strconv"
import "math/big"

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

func isPrime(num int64) bool {
	return big.NewInt(num).ProbablyPrime(100000)
}

func main() {
	runtime.GOMAXPROCS(1)

	pandigital := "987654321"

	for "" != pandigital {
		nines := permutations(pandigital)
		print(len(nines), "\n")
		for n := 0; n != len(nines); n++ {
			num := toInt(nines[n])
			if isPrime(num) {
				fmt.Printf("%v", num)
				return
			}
		}
		pandigital = pandigital[1:]
	}
	panic("None found")
}
