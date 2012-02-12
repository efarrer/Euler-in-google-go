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

main: main.go
	gofmt -w=true *.go && go build *.go

clean:
	rm main results
EOF_MAKE

git add $GOFILE
git add $MAKEFILE

$EDITOR $GOFILE
