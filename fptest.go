package main

import (
	"fmt"
	"os"
	"strconv"
)

func threshold(tValue float64) int64 {
	const k = 0x1p+56
	if tValue < 1.0 {
		return int64(k*tValue + 0.5)
	}
	return int64(k/tValue + 0.5)
}

func samplingRate(tValue float64) int64 {
	if tValue < 1.0 {
		return int64(1.0/tValue + 0.5)
	}
	return int64(tValue + 0.5)
}

func samplingProbability(tValue float64) float64 {
	if tValue < 1.0 {
		return tValue
	}
	return 1.0 / tValue
}

func all(args []string) {
	for _, arg := range args {
		tValue, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad value: %v\n", err)
			continue
		}
		fmt.Printf("tValue: %.14g\n", tValue)
		fmt.Printf("threshold: %d\n", threshold(tValue))
		fmt.Printf("samplingRate: %d\n", samplingRate(tValue))
		fmt.Printf("samplingProbability: %.14g\n", samplingProbability(tValue))
		fmt.Println()
	}
}

func thresh(arg string) string {
	tValue, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return "bad value: " + arg
	}
	return fmt.Sprintf("%d", threshold(tValue))
}

func rate(arg string) string {
	tValue, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return "bad value: " + arg
	}
	return fmt.Sprintf("%d", samplingRate(tValue))
}

func prob(arg string) string {
	tValue, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return "bad value: " + arg
	}
	return fmt.Sprintf("%.14g", samplingProbability(tValue))
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s command <tValue>...\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "commands: thresh, rate, prob, all\n")
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		usage()
	}
	switch args[0] {
	case "thresh":
		fmt.Println(thresh(args[1]))
	case "rate":
		fmt.Println(rate(args[1]))
	case "prob":
		fmt.Println(prob(args[1]))
	case "all":
		all(args[1:])
	default:
		usage()
	}
}
