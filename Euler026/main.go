package main

import (
	"fmt"
	"math/big"
	"runtime"
	"strings"
)

type cycle struct {
	denominator int
	fullString  string
	cycleChars  string
	cycleLength int
	cycleCount  int
}

func findCycles(denominator int, length int, str string) *cycle {
	if length == 1 {
		return nil
	}

	// Start at offset i
	for i := 0; i+length < len(str); i++ {
		// Starting at i get the first substring
		model := str[i : i+length]
		//fmt.Printf("Model %v\n", model)

		count := 0
		// All subsequent substrings must match the model substring
		for next := i + length; next < len(str); next += length {
			end := next + length
			if end > len(str) {
				end = len(str)
			}
			slice := str[next:end]

			truncModel := model
			if len(slice) < len(model) {
				truncModel = model[0:len(slice)]
			}
			//fmt.Printf("Length %v, Offset %v Start %v End %v Model %v String %v\n", length, i, next, end, truncModel, slice)
			if slice != truncModel {
				count = 0
				break
			}
			count += 1
		}

		// We found a match so return the result
		if count > 1 {
			other := findCycles(denominator, length-1, str)
			// If we got a response use the one that had the most cycles
			if nil != other && other.cycleCount > count {
				return other
			} else {
				return &cycle{denominator, str, model, length, count}
			}
		}
	}

	// Not a match so try a shorter string
	return findCycles(denominator, length-1, str)
}

func main() {
	runtime.GOMAXPROCS(1)

	biggest := &cycle{0, "", "", 0, 0}

	for d := int64(2); d < 1000; d++ {
		rat := big.NewRat(1, d)
		str := rat.FloatString(2001)
		// Split off the last character which may be rounded
		str = str[0 : len(str)-1]
		// Remove all before including "."
		str = strings.SplitAfter(str, ".")[1]
		// Remove all trailing 0's
		str = strings.TrimRight(str, "0")

		this := findCycles(int(d), 1000, str)

		if this != nil {
			if this.cycleLength > biggest.cycleLength {
				biggest = this
			}
		}
	}

	fmt.Printf("%v", biggest.denominator)
}
