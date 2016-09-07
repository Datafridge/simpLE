#!bin/sh
cd ~/projects/golang/src/github.com/mmbuw/simpLE 
go install
cd examples/
go build peripheral.go
./peripheral
