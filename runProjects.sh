#!/bin/bash -e

for dir in $(ls); do
    if [ -d "$dir" ]; then
        cd "$dir"
        gofmt -w=true *.go
        6g *.go
        6l -o main *.6
        echo -n "Euler 1: "
        ./main
        cd ../
    fi
done
