#!/bin/bash -e

for dir in $(ls); do
    if [ -d "$dir" ]; then
        cd "$dir"
        gofmt -w=true *.go
        6g *.go
        6l -o main *.6
        echo -n "${dir}: "
        ./main
        echo ""
        COUNT=$(echo "$dir" | sed 's/Euler//')
        ANSWER=$(./main)
        cd ../
        if ! grep "^${COUNT}[.] ${ANSWER}" ./EulerSolutions.txt > /dev/null; then
            echo "It appears that Euler #${COUNT} is wrong."
            exit 1
        fi
    fi
done
