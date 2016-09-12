package main

import (
  //"github.com/esiqveland/notify"
  "github.com/mmbuw/simpLE"
  "fmt"
)

var scanner simpLE.Scanner

func main()  {
  //var sender simpLE.Sender
  //sender.Start()

  scanner.Start(test_callback)
}

func test_callback() {
    results := scanner.Get_advertisements()
    //scanner.Get_advertisements()

    for key, _ := range results {
        fmt.Printf("Address: %v\n",key)
        fmt.Printf("Name: %v\n",results[key]["Name"])
        fmt.Printf("ManufacturerData: %v \n\n",results[key]["ManufacturerData"].(map[uint16][]uint8)[65535])
    }


    //fmt.Printf("%v",results["ManufacturerData"])
    //fmt.Printf("--------------------------------------------------------------")*/
}
