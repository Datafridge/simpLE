#!/usr/bin/env python3

import array
from datetime import datetime, date
import matplotlib.pyplot as plt

def main(tag):
    x = []
    actual = datetime(2010,1,1)
    end    = datetime(2016,7,30)
    total  = end - actual
    filename = tag+"_"+str(actual)+"_"+str(end)
    f = open(filename, 'r')
    for line in f:
        erg = line.split( )
        if len(x)==0:
            x = [int(erg[1])]
        else:
            x = x + [(x[-1]+int(erg[1]))]
    f.close()
    y = range(len(x))
    plot(x,y)


def plot(x,y):
    plt.plot(y, x, '-')
    plt.axis([min(y), max(y), min(x), max(x)])
    plt.show()

if __name__ == "__main__":
    main("bluetooth-lowenergy")
