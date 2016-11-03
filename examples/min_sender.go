package main

import (
  "github.com/mmbuw/simpLE"
)

var sender simpLE.Sender

func main()  {
  sender.Start("0123456789abcdef")
}
