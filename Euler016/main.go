package main

import "fmt"
import "runtime"
import "strconv"
import "math/big"

func main() {
	runtime.GOMAXPROCS(1)

	numStr := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1000), nil).String()

	total := 0
	for i := 0; i < len(numStr); i++ {
		res, _ := strconv.Atoi(numStr[i : i+1])
		total += res
	}
	fmt.Printf("%v", total)
}
