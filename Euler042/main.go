package main

import "fmt"
import "runtime"
import "os"
import "io"
import "bufio"
import "strings"

func readAll(file string) string {
	res := ""
	fd, err := os.Open(file)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Can't open words: "+err.Error())
		os.Exit(1)
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	for {
		this, err := reader.ReadString(0)
		if err == io.EOF {
			res += this
			break
		}
		if nil != err {
			fmt.Fprintf(os.Stderr, "Can't read from words: "+err.Error())
			os.Exit(1)
		}
		res += this
	}

	return res
}

func main() {
	runtime.GOMAXPROCS(1)

	file := readAll("./words.txt")
	file = strings.Replace(file, "\"", "", -1)
	words := strings.Split(file, ",")

	// Figure out the max triangle word value that needs to be calculated
	maxPossibleTriangleValue := 0
	for _, word := range words {
		thisMax := len(word) * 26
		if thisMax > maxPossibleTriangleValue {
			maxPossibleTriangleValue = thisMax
		}
	}

	// Calculate all the triangle words that we care about
	triangle := make(map[int]bool)
	for n := 0; true; n++ {
		lastTriangle := (n * (n + 1)) / 2
		triangle[lastTriangle] = true
		if lastTriangle > maxPossibleTriangleValue {
			break
		}
	}

	// Iterate through the words and calculate their value
	count := 0
	for _, word := range words {
		wordValue := 0
		for _, letter := range word {
			wordValue += 1 + (int(letter) - 'A')
		}
		if triangle[wordValue] {
			count++
		}
	}

	fmt.Printf("%v", count)
}
