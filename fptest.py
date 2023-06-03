#!/usr/bin/env python3

import sys
import os
import random

def threshold(tValue):
    k = 0x100_0000_0000_0000  # 2^56
    if tValue < 1.0:
        return int(k * tValue + 0.5)
    else:
        return int(k / tValue + 0.5)

def samplingRate(tValue):
    if tValue < 1.0:
        return int(1.0/tValue + 0.5)
    else:
        return int(tValue + 0.5)

def samplingProbability(tValue):
    if tValue < 1.0:
        return tValue
    else:
        return 1.0/tValue

def all(args):
    for i in args:
        tValue = float(i)

        print(f"tValue: {tValue:.14g}")
        print(f"threshold: {threshold(tValue):d}")
        print(f"samplingRate: {samplingRate(tValue):d}")
        print(f"samplingProbability: {samplingProbability(tValue):.14g}")
        print()

def thresh(arg):
    tValue = float(arg)
    return f"{threshold(tValue):d}"

def rate(arg):
    tValue = float(arg)
    return f"{samplingRate(tValue):d}"

def prob(arg):
    tValue = float(arg)
    return f"{samplingProbability(tValue):.14g}"


if __name__ == '__main__':
    if len(sys.argv) == 1:
        print("Usage: fptest.py command <tValue> ...")
        print("command: all, thresh, rate, prob")
        sys.exit(1)

    args = sys.argv[2:]
    command = sys.argv[1]
    if command == "all":
        all(args)
    elif command == "thresh":
        print(thresh(args[0]))
    elif command == "rate":
        print(rate(args[0]))
    elif command == "prob":
        print(prob(args[0]))
    else:
        print("Usage: fptest.py command <tValue> ...")
        print("command: all, thresh, rate, prob")
        sys.exit(1)

