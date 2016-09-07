package simpLE

import (
    "github.com/godbus/dbus"
    //"errors"
    "fmt"
    //"reflect"
)

type Scanner struct {
    dev *device
    res map[string]remote_device
    f   func()
}

func (s *Scanner) Start() error {
    s.res = make(map[string]remote_device)

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

		fmt.Printf("Sender: %v\nPath: %v\nName: %T\nBody:%v\n\n",v.Sender,v.Path,v.Name,v.Body[0])


        if v.Sender == ":1.3" {
            fmt.Printf("sender is bluez \n")
            switch v.Name {
            case "org.freedesktop.DBus.ObjectManager.InterfacesAdded":
                //TODO interface hinzufügen
                fmt.Printf("interface was added \n")

                index := string(v.Body[0].(dbus.ObjectPath))
                fmt.Printf("index type: %T \n", index)

                s.res[index] = remote_device{}
                rdt := s.res[index]
                rd  := &rdt
                rd.set_path(v.Body[0].(dbus.ObjectPath))
                fmt.Printf("new remote_device created\n")
                fmt.Printf("path type: %T value: %v \n",string(rd.get_path()),string(rd.get_path()))

            case "org.freedesktop.DBus.ObjectManager.InterfacesRemoved":
                //TODO interface aus Ergebnissen löschen
                fmt.Printf("interface was removed \n")
            }
        }

	}
    return nil
}

func (s *Scanner) save(sig *dbus.Signal) error {
    // check if the sender is bluez
    if sig.Sender == ":1.3" {

    }


    return nil
}
