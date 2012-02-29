package main

import "fmt"
import "runtime"
import "strconv"

func toInt(num string) int64 {
	i, e := strconv.ParseInt(num, 10, 64)
	if nil != e {
		panic(e)
	}
	return i
}

func main() {
	runtime.GOMAXPROCS(1)

	print("start\n")
	str := ""
	for i := int64(1); len(str) < 1000000; i++ {
		str += strconv.FormatInt(i, 10)
	}
	print("finish generating string\n")

	res := int64(1)
	for i := int64(1); i <= 1000000; i *= 10 {
		print("i=", i, "\n")
		res *= toInt(str[i-1 : i])
	}
	print("done\n")
	fmt.Printf("%v", res)
}
