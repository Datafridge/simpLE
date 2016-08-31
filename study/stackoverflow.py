#!/usr/bin/env python3

from datetime import datetime, date, timedelta
import requests
import json

def main(tag):
    actual = datetime(2010,1,1)
    end    = datetime(2016,7,30)
    total  = end - actual
    filename = tag+"_"+str(actual)+"_"+str(end)
    f = open(filename, 'w')
    f.close()


    while True:
        f = open(filename, 'a')
        d = end - actual
        f.write(get_data(tag,actual))
        f.close()
        print("progress:"+str(total.days-d.days)+"/"+str(total.days)+" "+str(quota_remaining)+" requests left", end="\r")
        if actual == end:
            break
        actual = actual+timedelta(1)

def get_data(tag,actual):
    entry = create_request(int(actual.timestamp()),int((actual+timedelta(1)).timestamp()),tag)
    global quota_remaining
    quota_remaining = entry["quota_remaining"]
    #print  (str(actual.year)+"-"+str(actual.month)+"-"+str(actual.day)+": "+str(entry["total"])+"\n")
    return (str(actual.year)+"-"+str(actual.month)+"-"+str(actual.day)+": "+str(entry["total"])+"\n")

def create_request(from_date,to_date,tag):
    r = requests.get("https://api.stackexchange.com/2.2/search?fromdate="+str(from_date)+"&todate="+str(to_date)+"&order=desc&sort=creation&tagged="+tag+"&site=stackoverflow&filter=!.UE8F0_CXHtXgLFC&key=maUPfGS71Y2I2AkkdnMpVw((")
    #print(r.json())
    return(r.json())

if __name__ == "__main__":
    main("ble")
