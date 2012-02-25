#!/bin/bash -e

for dir in $(ls); do
    if [ -d "$dir" ]; then
        make -C "$dir" > /dev/null
        ANSWER=$(cat "${dir}/results")
        echo "${dir}: $ANSWER"
        COUNT=$(echo "$dir" | sed 's/Euler//')
        if ! grep "^${COUNT}[.] ${ANSWER}" ./EulerSolutions.txt > /dev/null; then
            echo "It appears that Euler #${COUNT} is wrong."
            exit 1
        fi
    fi
done
