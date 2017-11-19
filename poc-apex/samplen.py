#!/usr/bin/env python2

# Sample X lines from a wordlist in an efficient manner.
# From: http://data-analytics-tools.blogspot.com.au/2009/09/reservoir-sampling-algorithm-in-perl.html

import sys
import random

if len(sys.argv) == 3:
    input = open(sys.argv[2],'r')
elif len(sys.argv) == 2:
    input = sys.stdin;
else:
    sys.exit("Usage: python samplen.py <lines> <?file>")

N = int(sys.argv[1]);
sample = [];

for i,line in enumerate(input):
    if i < N:
        sample.append(line)
    elif i >= N and random.random() < N/float(i+1):
        replace = random.randint(0,len(sample)-1)
        sample[replace] = line

for line in sample:
    sys.stdout.write(line)
