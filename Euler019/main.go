package main

import "fmt"
import "runtime"
import "time"

func main() {
	runtime.GOMAXPROCS(1)

	now := time.Now()

	// Walk back to the last day of the twentieth century
	for now.Year() != 2000 || now.Month() != time.December || now.Day() != 31 {
		now = now.AddDate(0, 0, -1)
	}

	count := 0
	for now.Year() > 1900 {
		if now.Day() == 1 {
			if now.Weekday() == time.Sunday {
				count++
			}
		}
		now = now.AddDate(0, 0, -1)
	}

	fmt.Printf("%v", count)
}
