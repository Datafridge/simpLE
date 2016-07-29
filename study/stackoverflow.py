#!/usr/bin/env python3

from datetime import datetime, date, timedelta
import requests

def main(tag):
    actual = datetime(2016,7,25)
    end    = datetime(2016,7,30)
    total  = end - actual
    f = open(tag+"_"+str(actual)+"_"+str(end), 'w')

    while True:
        d = end - actual
        print("progress:"+str(total.days-d.days)+"/"+str(total.days), end="\r")
        f.write(get_data(tag,actual))
        if actual == end:
            break
        actual = actual+timedelta(1)

def get_data(tag,actual):
    return str(create_request(int(actual.timestamp()),int((actual+timedelta(1)).timestamp()),tag))
    #TODO Format JSON


def create_request(from_date,to_date,tag):
    r = requests.get("https://api.stackexchange.com/2.2/search?fromdate="+str(from_date)+"&todate="+str(to_date)+"&order=desc&sort=creation&tagged="+tag+"&site=stackoverflow&filter=!.UE8F0_CXHtXgLFC&key=maUPfGS71Y2I2AkkdnMpVw((")
    #print(r.json())
    return(r.json())

if __name__ == "__main__":
    main("bluetooth-lowenergy")
