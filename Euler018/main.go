package main

import "fmt"
import "runtime"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

type node struct {
	value int // The value of the node
	max   int // The maximum path value of the node
}

func readTriangle() [][]*node {
	tri := [][]*node{}

	fd, err := os.Open("./triangle.txt")
	if nil != err {
		fmt.Fprintf(os.Stderr, "Can't open triangle")
		os.Exit(1)
	}
	defer fd.Close()

	lineReader := bufio.NewReader(fd)

	for {
		bytes, isPrefix, err := lineReader.ReadLine()
		if nil != err {
			// We must be at the end so return
			return tri
		}
		line := string(bytes)
		for isPrefix {
			bytes, isPrefix, err = lineReader.ReadLine()
			if nil != err {
				fmt.Fprintf(os.Stderr, "Can't read rest of line")
				os.Exit(1)
			}
			line += string(bytes)
		}

		row := []*node{}

		numbers := strings.Split(line, " ")
		for i := 0; i < len(numbers); i++ {
			number, err := strconv.Atoi(numbers[i])
			if nil != err {
				fmt.Fprintf(os.Stderr, "Can't convert number %v\n", numbers[i])
				os.Exit(1)
			}
			row = append(row, &node{number, 0})
		}

		tri = append(tri, row)
	}

	// Should never be called because of the forever loop above
	return nil
}

func main() {
	runtime.GOMAXPROCS(1)

	triangle := readTriangle()

	for h := 0; h < len(triangle); h++ {
		for w := 0; w < len(triangle[h]); w++ {
			current := triangle[h][w]

			parent0Max := 0
			parent1Max := 0
			// The top has no parents
			if h != 0 {
				// If we're on the left edge then we only have one parent
				if w != 0 {
					parent0Max = triangle[h-1][w-1].max
				}
				// If we're on the right edge then we only have one parent
				if w < len(triangle[h-1]) {
					parent1Max = triangle[h-1][w].max
				}
			}
			// The max path is itself + the max(maxpath(parent0), maxpath(parent1))
			current.max = current.value + int(math.Max(float64(parent0Max), float64(parent1Max)))
		}
	}

	lastrow := len(triangle) - 1
	biggestMax := 0
	for w := 0; w < len(triangle[lastrow]); w++ {
		biggestMax = int(math.Max(float64(biggestMax), float64(triangle[lastrow][w].max)))
	}
	fmt.Printf("%v", biggestMax)
}
