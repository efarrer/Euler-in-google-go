package main

import "fmt"
import "runtime"
import "os"
import "io/ioutil"
import "strings"
import "sort"

func loadNames() []string {
	fd, err := os.Open("names.txt")
	if nil != err {
		fmt.Fprintf(os.Stderr, "Can't open names")
		os.Exit(1)
	}
	defer fd.Close()

	nameBytes, err := ioutil.ReadAll(fd)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Can't read names")
		os.Exit(1)
	}

	names := strings.Replace(string(nameBytes), "\"", "", -1)

	return strings.Split(names, ",")
}

func sum(nums []int) int {
	total := 0
	for i := 0; i != len(nums); i++ {
		total += nums[i]
	}
	return total
}

func charsToInt(str string) []int {
	ints := make([]int, len(str))

	for i := 0; i < len(str); i++ {
		ints[i] = 1 + (int(str[i]) - int('A'))
	}

	return ints
}

func main() {
	runtime.GOMAXPROCS(1)

	names := loadNames()
	sort.Sort(sort.StringSlice(names))

	total := 0

	for i := 0; i < len(names); i++ {
		total += (i + 1) * sum(charsToInt(names[i]))
	}
	fmt.Printf("%v", total)
}
