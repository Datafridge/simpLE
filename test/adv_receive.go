package main

import (
	//"encoding/json"
	"github.com/godbus/dbus"
	//"github.com/godbus/dbus/prop"
	//"github.com/godbus/dbus/introspect"
	//"os"
	//"time"
	"fmt"
)

type foo string

func main() {
	//connect to systembus
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}

	adapter_path := dbus.ObjectPath("/org/bluez/hci0")
	adapter := conn.Object("org.bluez",adapter_path)

	var result interface{}
	err = adapter.Call("org.bluez.Adapter1.StartDiscovery", 0).Store(&result)

	conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='/',interface='org.freedesktop.DBus.ObjectManager',sender='org.bluez'")

	conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='/',interface='org.freedesktop.DBus.Properties',sender='org.bluez'")

	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)
	for v := range c {
		fmt.Printf("Sender: %v\nPath: %v\nName: %v\nBody:%v\n\n",v.Sender,v.Path,v.Name,v.Body[0])
	}
}
