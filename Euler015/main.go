package main

import "fmt"
import "runtime"

// Threat the boxes like a graph of n+1 nodes. Then count the unique ways to
// traverse the graph
const size = 20

const width = size + 1
const height = size + 1

const startx = 0
const starty = 0

const endx = width - 1
const endy = width - 1

func main() {
	runtime.GOMAXPROCS(1)

	const empty = -1

	grid := make([][]int64, width)
	for x := 0; x != width; x++ {
		grid[x] = make([]int64, height)
	}

	for x := 0; x != width; x++ {
		for y := 0; y != height; y++ {
			grid[x][y] = empty
		}
	}

	// Work back to the beginning and calculate the path as we go
	for x := endx; x >= startx; x-- {
		for y := endy; y >= starty; y-- {
			paths := int64(0)

			// Set the path to the end
			if x == endx && y == endy {
				paths = int64(1)
			}

			if x+1 < width {
				paths += grid[x+1][y]
			}
			if y+1 < height {
				paths += grid[x][y+1]
			}
			grid[x][y] = paths
		}
	}

	fmt.Printf("%v", grid[startx][starty])
}
