#!/usr/bin/python

import numpy as np
import matplotlib.pyplot as plt
import random 
import sys
import math

N = 100
AMPLITUDE = 4
MIN_DILATION = 1
MAX_DILATION = 10
MIN_PHASE = -math.pi
MAX_PHASE = math.pi
MAPWIDTH = 4000

def surfacefun(x, param):
    result = 0
    for (a,p,d) in param:
        result = result + a * np.sin(0.01*math.log(d)*(x-p))
        #to make shure there arent any impossible jumps... 
    return result


def gen_parameters(n):
    result = []
    for i in range(n):
        a = random.randint(-AMPLITUDE,AMPLITUDE)
        p = random.randint(int(MIN_PHASE*100), int(MAX_PHASE*100)) / 100.0
        d = random.randint(MIN_DILATION, MAX_DILATION)
        result.append((a,p,d))
    return result

if len(sys.argv)>1:
    print("seeding rng with " + sys.argv[1])
    random.seed(sys.argv[1])
else:
    print("seeding rng with system time")
    random.seed
 
param = gen_parameters(N)

x = np.arange(MAPWIDTH)
y = np.vectorize(lambda x : int(surfacefun(x,param)))(x)
for index, current in enumerate(y):
    if (index>0):
        previous = y[index-1]
        if previous+3 < current:
            print("too steep at " + str(index))
            y[index] = previous+3
        if previous-3 > current:
            print("too steep at " + str(index))
            y[index] = previous-3

line, = plt.plot(x, y, '-', linewidth = 2) 

plt.show()
