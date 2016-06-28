package main

import (
  "runtime"
  "fmt"
  "golang.org/x/mobile/app"
)

func main (){
  app.Main(func(a app.App) {
    fmt.Println(runtime.GOOS)
	})
}
