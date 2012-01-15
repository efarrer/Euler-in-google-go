package main

import "fmt"
import "runtime"
import "strconv"
import "math/big"

func main() {
	runtime.GOMAXPROCS(1)

	fact := big.NewInt(int64(1))
	for i := int64(1); i <= int64(100); i++ {
		fact.Mul(fact, big.NewInt(i))
	}

	factStr := fact.String()

	sum := 0
	for i := 0; i < len(factStr); i++ {
		curr, _ := strconv.Atoi(factStr[i : i+1])
		sum += curr
	}

	fmt.Printf("%v", sum)
}
