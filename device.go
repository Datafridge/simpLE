package simpLE

import (
    "encoding/json"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
)

type device struct {
    path     string
    uuids    []string
    address  string
    alias    string
    modalias string
    name     string
    class    uint32
}

type local_device struct {
    device
    le                  bool
    discoverable        bool
    discovering         bool
    pairable            bool
    powered             bool
    discoverableTimeout uint32
    pairableTimeout     uint32
}

type remote_device struct {
    device
    blocked          bool
    connected        bool
    legacyPairing    bool
    paired           bool
    servicesResolved bool
    trusted          bool
    serviceData      map[string][]uint8
    manufacturerData map[uint16][]uint8
    rssi             int16
    txpower          int16
    adapter          string
    icon             string
    appearance       uint16
}

func (d *device) dbusobject()  dbus.BusObject{

    bus, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}

    adapter_path := dbus.ObjectPath("/org/bluez/hci0")
	adapter := bus.Object("org.bluez",adapter_path)
    return adapter
}

type temp_Device struct {
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
	var device temp_Device

	json.Unmarshal(data, &device)
	return device.Children[0].Name
}
