package main

import (
  "github.com/esiqveland/notify"
  "github.com/mmbuw/simpLE"
  "fmt"
)

func main()  {
  //var sender simpLE.Sender
  //sender.Start()

  var scanner simpLE.Scanner
  scanner.Start(test_callback)
}

func test_callback() {
    fmt.Printf("Advertisement erhalten")
}
