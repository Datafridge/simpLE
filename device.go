package simpLE

import (
    "github.com/godbus/dbus"
)

type device struct {
    path string
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
