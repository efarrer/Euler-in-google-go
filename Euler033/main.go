package main

import "fmt"
import "runtime"
import "math/big"
import "strconv"

func main() {
	runtime.GOMAXPROCS(1)

	prod := big.NewRat(1, 1)
	for num := int64(10); num < 99; num++ {
		for den := int64(num + 1); den < 100; den++ {
			numStr := strconv.FormatInt(num, 10)
			denStr := strconv.FormatInt(den, 10)

			numFirst, _ := strconv.ParseInt(numStr[:1], 10, 64)
			numSecond, _ := strconv.ParseInt(numStr[1:], 10, 64)
			denFirst, _ := strconv.ParseInt(denStr[:1], 10, 64)
			denSecond, _ := strconv.ParseInt(denStr[1:], 10, 64)

			// Do the stupid math by removing a matching type
			rat := big.NewRat(num, den)
			if 0 != numFirst && 0 != denSecond && numFirst == denFirst {
				if big.NewRat(numSecond, denSecond).Cmp(rat) == 0 {
					prod.Mul(prod, rat)
				}
			}
			if 0 != numFirst && 0 != denFirst && numFirst == denSecond {
				if big.NewRat(numSecond, denFirst).Cmp(rat) == 0 {
					prod.Mul(prod, rat)
				}
			}
			if 0 != numSecond && 0 != denSecond && numSecond == denFirst {
				if big.NewRat(numFirst, denSecond).Cmp(rat) == 0 {
					prod.Mul(prod, rat)
				}
			}
			if 0 != numSecond && 0 != denFirst && numSecond == denSecond {
				if big.NewRat(numFirst, denFirst).Cmp(rat) == 0 {
					prod.Mul(prod, rat)
				}
			}
		}
	}

	fmt.Printf("%v", prod.Denom())
}
