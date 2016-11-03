package main

import (
  "github.com/mmbuw/simpLE"
  "fmt"
)

var scanner simpLE.Scanner
var sender simpLE.Sender

func main()  {
  scanner.Start(print_information)
  go sender.Start()
}

func print_information() {
    results := scanner.Get_last_advertisements()

    fmt.Printf("Address: %v\n",results["Address"])
    fmt.Printf("Name: %v\n",results["Name"])
    fmt.Printf("ManufacturerData: %v \n\n",results["ManufacturerData"].(map[uint16][]uint8)[65535])
}
