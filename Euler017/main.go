package main

import "fmt"
import "runtime"
import "strings"

var underTwentyMap map[int]string
var tensMap map[int]string

const hundred = "hundred"
const thousand = "thousand"

func init() {
	underTwentyMap = map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen"}
	tensMap = map[int]string{10: "ten", 20: "twenty", 30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty", 90: "ninety"}
}

func toString(num int) string {
	result := ""

	thousands := num / 1000
	num -= (1000 * thousands)
	if thousands > 0 {
		result += underTwentyMap[thousands] + " " + thousand + " "
	}

	hundreds := num / 100
	num -= (100 * hundreds)
	if hundreds > 0 {
		result += underTwentyMap[hundreds] + " " + hundred + " "
	}

	if 0 == num {
		return result
	} else if result != "" {
		result += "and "
	}

	if num < 20 {
		result += underTwentyMap[num]
	} else {
		tens := num / 10
		num -= (10 * tens)
		result += tensMap[tens*10] + "-" + underTwentyMap[num]
	}

	return result
}

func main() {
	runtime.GOMAXPROCS(1)

	results := ""
	for i := 1; i <= 1000; i++ {
		results += toString(i)
	}
	stripped := strings.Replace(strings.Replace(results, " ", "", -1), "-", "", -1)
	fmt.Printf("%v", len(stripped))
}
