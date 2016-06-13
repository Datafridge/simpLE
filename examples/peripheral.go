package main

import (
  "github.com/mmbuw/simpLE"
  "fmt"
)

func main()  {
  dev := simpLE.Get_devices()
  fmt.Println(len(dev))
}
