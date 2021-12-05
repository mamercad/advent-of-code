#!/usr/bin/env python3

with open("aoc01.dat") as f:
    lines = f.read().splitlines()

increases = 0
for i in range(len(lines)-1):
    if int(lines[i]) < int(lines[i+1]):
        increases += 1

print(increases)