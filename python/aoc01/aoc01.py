#!/usr/bin/env python3

with open("aoc01.dat") as f:
    lines = f.read().splitlines()

# part 1

increases = 0
for i in range(len(lines) - 1):
    if int(lines[i]) < int(lines[i + 1]):
        increases += 1

print(f"part 1 increases: {increases}")

# part 2

increases = 0
for i in range(len(lines) - 3):
    window1 = int(lines[i]) + int(lines[i + 1]) + int(lines[i + 2])
    window2 = int(lines[i + 1]) + int(lines[i + 2]) + int(lines[i + 3])
    if window1 < window2:
        increases += 1

print(f"part 2 increases: {increases}")
