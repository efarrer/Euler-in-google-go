package main

import "fmt"
import "runtime"

const max = 1000000

var cache map[int64]int64

func init() {
	cache = make(map[int64]int64)
}

func cacheAndReturn(num, colltz int64) int64 {
	if num < max {
		cache[num] = colltz
	}
	return colltz
}

func lookupColltz(num int64) (int64, bool) {
	colltz, ok := cache[num]
	if ok {
	}
	return colltz, ok
}

func calculateColltz(num int64) int64 {

	// See if we have it cached already
	colltz, ok := lookupColltz(num)
	if ok {
		return colltz
	}

	if 1 == num {
		return cacheAndReturn(num, 1)
	}

	next := int64(0)
	if int64(0) == (num % int64(2)) {
		next = num / int64(2)
	} else {
		next = (int64(3) * num) + int64(1)
	}

	return cacheAndReturn(num, 1+calculateColltz(next))
}

func main() {
	runtime.GOMAXPROCS(1)

	for i := int64(1); i != max; i++ {
		calculateColltz(i)
	}

	maxKey, maxValue := int64(0), int64(0)
	for key, value := range cache {
		if value > maxValue {
			maxKey = key
			maxValue = value
		}
	}

	fmt.Printf("%v\n", maxKey)
}
