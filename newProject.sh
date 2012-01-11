#!/bin/bash -e

PROBLEM=$(printf "%03d" "$1")

if [ "" = "$1" ]; then
    echo "Usage $0 <problem_number>"
    exit 1
fi



DIR="Euler${PROBLEM}"
MAKEFILE="${DIR}/Makefile"
GOFILE="${DIR}/main.go"

if [ -d "$DIR" ]; then
    echo "Problem already exists"
    exit 1
fi

mkdir $DIR

cat >$GOFILE <<EOF_GO
package main

import "fmt"
import "runtime"

func main() {
    runtime.GOMAXPROCS(1)

    fmt.Printf("TODO $PROBLEM")
}
EOF_GO

gofmt -w=true $GOFILE

cat >$MAKEFILE <<EOF_MAKE
results: main
	./main > results

main: main.6
	6l -o main *.6

main.6: main.go
	gofmt -w=true *.go && 6g *.go

clean:
	rm main *.6 results
EOF_MAKE

git add $GOFILE
git add $MAKEFILE

$EDITOR $GOFILE
