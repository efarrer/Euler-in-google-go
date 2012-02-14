package main

import "fmt"
import "runtime"
import "math/big"
import "sort"

type BigIntSlice []*big.Int

func (is BigIntSlice) Len() int {
	return len(is)
}

func (is BigIntSlice) Less(i, j int) bool {
	return is[i].Cmp(is[j]) == -1
}

func (is BigIntSlice) Swap(i, j int) {
	t := is[i]
	is[i] = is[j]
	is[j] = t
}

func main() {
	runtime.GOMAXPROCS(1)

	const min = int64(2)
	const max = int64(100)

	results := make([]*big.Int, 0, 1000)
	results = append(results, big.NewInt(4))
	results = append(results, big.NewInt(4))
	results = append(results, big.NewInt(4))
	results = append(results, big.NewInt(4))
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			result := big.NewInt(0)
			result.Exp(big.NewInt(a), big.NewInt(b), nil)
			results = append(results, result)
		}
	}

	sort.Sort(BigIntSlice(results))

	pruned := make([]*big.Int, len(results))
	var last *big.Int = nil
	p := 0
	for r := 0; r < len(results); r++ {
		if last == nil || results[r].Cmp(last) != 0 {
			pruned[p] = results[r]
			last = results[r]
			p++
		}
	}
	fmt.Printf("%v", len(pruned[0:p]))
}
