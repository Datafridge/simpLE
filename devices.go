package simpLE

import (
	"encoding/json"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
	//"os"
	//"fmt"
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

// get all available devices
func Get_devices() string{
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	node, err := introspect.Call(conn.Object("org.bluez", "/org/bluez"))
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(node, "", "    ")
	var device Device

	json.Unmarshal(data, &device)
	return device.Children[0].Name
}
