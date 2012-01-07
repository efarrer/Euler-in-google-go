#!/bin/bash -e

PROBLEM="$1"

if [ "" = "$1" ]; then
    echo "Usage $0 <problem_number>"
    exit 1
fi



DIR="Euler${PROBLEM}"
FILE="${DIR}/main.go"

if [ -d "$DIR" ]; then
    echo "Problem already exists"
    exit 1
fi

mkdir $DIR

cat >$FILE <<EOF
package main

import "fmt"

func main() {
    fmt.Printf("TODO $PROBLEM")
}
EOF

gofmt -w=true $FILE

