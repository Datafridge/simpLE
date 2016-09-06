package simpLE

import (
    "github.com/godbus/dbus"
    //"errors"
    //"fmt"
)

type Scanner struct {
    dev *device
    res map[string]advertisement
    f   func()
}

func (s *Scanner) Start() error {

    bus, err := dbus.SystemBus()
    if err != nil {
        panic(err)
    }

    var dev device
    dev.dbusobject().Call("org.bluez.Adapter1.StartDiscovery", 0).Store(&result)

    bus.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='/',interface='org.freedesktop.DBus.ObjectManager',sender='org.bluez'")

	bus.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='/',interface='org.freedesktop.DBus.Properties',sender='org.bluez'")

	c := make(chan *dbus.Signal, 10)
	bus.Signal(c)
	for v := range c {
        //objects := v.Body[1].(map[string]map[string]dbus.Variant)
        //fmt.Printf("test")
		//fmt.Printf("Sender: %v\nPath: %v\nName: %v\nBody:%v\n\n",v.Sender,v.Path,v.Name,test[0])
        //fmt.Printf("Typ: %T\nValue:%v\n\n",objects["org.bluez.Device1"]["Class"],objects["org.bluez.Device1"]["Class"])
	}

    return nil
}
