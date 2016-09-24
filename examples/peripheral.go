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
    results := scanner.Get_last_advertisements()

    fmt.Printf("Address: %v\n",results["Address"])
    fmt.Printf("Name: %v\n",results["Name"])
    fmt.Printf("ManufacturerData: %v \n\n",results["ManufacturerData"].(map[uint16][]uint8)[65535])
}
