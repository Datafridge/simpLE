package main

import (
	"encoding/json"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
	//"os"
	"fmt"
)

type Device struct {
    XMLName string
		Name string
		Interfaces interface {}
		Children []struct {
			XMLName interface {}
			Name string
		}
}

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	node, err := introspect.Call(conn.Object("org.bluez", "/org/bluez"))
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(node, "", "    ")
	//os.Stdout.Write(data)
	get_devices(data)
}

func get_devices(introspect_json []byte)  {
	var device Device
	//var parsed map[string]interface{}
	json.Unmarshal(introspect_json, &device)
	fmt.Printf(device.Children[0].Name)
}
