#! /usr/bin/env bash

compare() {
    g=`go run fptest.go thresh $1`
    p=`python3 fptest.py thresh $1`
    j=`node fptest.js thresh $1`

    if [ $g == $p ] && [ $p == $j ]; then
        echo "tValue $1 gives threshold $g"
    else
        echo "No match for $1!"
        echo "        Go: $g"
        echo "    Python: $p"
        echo "Javascript: $j"
    fi
}

for i in $*; do
    compare $i
done
