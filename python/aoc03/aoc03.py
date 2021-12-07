#!/usr/bin/env python3

def read_diagnostic_report(filename: str):
    with open(filename) as f:
        lines = f.read().splitlines()
    d = []
    for line in lines:
        d.append(list(line))
    return d


def slice_column(matrix: list, column: int):
    col = []
    for row in matrix:
        col.append(row[column])
    return col


def get_frequency(l: list):
    freq = dict()
    for i in l:
        if i in freq:
            freq[i] += 1
        else:
            freq[i] = 1
    return freq


def most_frequent(d: dict):
    return max(d, key=d.get)


def least_frequent(d: dict):
    return min(d, key=d.get)


def bin_to_dec(l: list):
    dec = 0
    for i in range(len(l)):
        dec += 2**i * int(l[-1*(i+1)])
    return dec


report = read_diagnostic_report("aoc03.dat")
columns = len(report[0])

gamma_rate, epsilon_rate = [], []
for column in range(columns):
    c = slice_column(report, column)
    f = get_frequency(slice_column(report, column))
    gamma_rate.append(most_frequent(f))
    epsilon_rate.append(least_frequent(f))

gr = bin_to_dec(gamma_rate)
er = bin_to_dec(epsilon_rate)

print(f"gamma rate: {gr}")
print(f"epsilon rate: {er}")
print(f"product: {gr * er}")
