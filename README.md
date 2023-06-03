# tvalue

This repository contains 3 different demonstration implementations (in Go, Python, and JavaScript) of calculations from tValue to threshold, sample rate, and sampling probability.

The point was to show the math, and in particular the rounding, across several languages.

This is done as commentary to https://github.com/open-telemetry/oteps/pull/226.

There is a small bash shell script to compare the output from the three programs.

To run all the code, you need to have all of these installed:

- Go
- Node
- Python3

None of them use any special libraries.

## syntax
All 3 programs have the same basic syntax:

```bash
<invoke> command <tValue>...
```

## running the programs
`<invoke>` can be:
- `go run fptest.go`
- `python3 fptest.py`
- `node fptest.js`

## commands
- `thresh` -- calculate the threshold value out of 2**56
- `rate` -- calculate the integer sample rate (the value of N for 1 in N),
- `prob` -- calculate sampling probability from 0-1
- `all` -- print a formatted list of all these values

