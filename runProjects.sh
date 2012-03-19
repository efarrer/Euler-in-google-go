#!/bin/bash -e

for dir in $(ls); do
    if [ -d "$dir" ]; then
        make -C "$dir" > /dev/null
        ANSWER=$(cat "${dir}/results")
        echo "${dir}: $ANSWER"
        COUNT=$(echo "$dir" | sed 's/Euler//')
        if ! grep "^${COUNT}[.] ${ANSWER}" ./EulerSolutions.txt > /dev/null; then
            echo "It appears that Euler #${COUNT} is wrong."
            echo "Got \"${COUNT}. $ANSWER\""
            expecting=$(grep "^${COUNT}[.].*" ./EulerSolutions.txt | tr -d '\r')
            echo "Expecting \"$expecting\""
            exit 1
        fi
    fi
done
