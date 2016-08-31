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

var BLUEZ_SERVICE_NAME = "org.bluez"
var LE_ADVERTISING_MANAGER_IFACE = "org.bluez.LEAdvertisingManager1"
var DBUS_OM_IFACE = "org.freedesktop.DBus.ObjectManager"
var DBUS_PROP_IFACE = "org.freedesktop.DBus.Properties"

var LE_ADVERTISEMENT_IFACE = "org.bluez.LEAdvertisement1"

type foo string

func main() {
	//connect to systembus
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}

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
