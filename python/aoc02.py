#!/usr/bin/env python3

with open("aoc02.dat") as f:
    lines = f.read().splitlines()

position = [0, 0]
for line in lines:
    direction, distance = line.split()
    if direction == "forward":
        position[0] += int(distance)
    elif direction == "up":
        position[1] -= int(distance)
    elif direction == "down":
        position[1] += int(distance)
    else:
        raise Exception(f"unknown direction: {direction}")

print(f"position: {position}")
print(f"product: {position[0] * position[1]}")
