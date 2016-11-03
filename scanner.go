package simpLE

import (
    "github.com/godbus/dbus"
    //"fmt"
)

type Scanner struct {
    dev  *device
    res  map[string]remote_device
    f    func()
    last remote_device
}

func (s *Scanner) Start(f1 func()) error {
    s.f = f1
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

		//fmt.Printf("Sender: %v\nPath: %v\nName: %T\nBody:%v\n\n",v.Sender,v.Path,v.Name,v.Body[0])



        if v.Sender == ":1.7" {
            //fmt.Printf("sender is bluez \n")
            switch v.Name {
            case "org.freedesktop.DBus.ObjectManager.InterfacesAdded":
                //fmt.Printf("interface was added \n")

                index := string(v.Body[0].(dbus.ObjectPath))
                //fmt.Printf("index type: %T \n", index)

                s.res[index] = remote_device{}
                rdt := s.res[index]
                rd  := &rdt
                //fmt.Printf("new remote_device created\n")
                rd.set_path(v.Body[0].(dbus.ObjectPath))

                objects1 := v.Body[1].(map[string]map[string]dbus.Variant)
                objects := objects1["org.bluez.Device1"]

                if objects["UUIDs"].Value() != nil {
                    rd.set_uuids(objects["UUIDs"].Value().([]string))
                }

                if objects["Address"].Value() != nil {
                    rd.set_address(objects["Address"].Value().(string))
                }

                if objects["Address"].Value() != nil {
                    rd.set_alias(objects["Alias"].Value().(string))
                }

                if objects["Modalias"].Value() != nil {
                    rd.set_modalias(objects["Modalias"].Value().(string))
                }

                if objects["Name"].Value() != nil {
                    rd.set_name(objects["Name"].Value().(string))
                }

                if objects["Class"].Value() != nil {
                    rd.set_class(objects["Class"].Value().(uint32))
                }

                if objects["Blocked"].Value() != nil {
                    rd.set_blocked(objects["Blocked"].Value().(bool))
                }

                if objects["Connected"].Value() != nil {
                    rd.set_connected(objects["Connected"].Value().(bool))
                }

                if objects["LegacyPairing"].Value() != nil {
                    rd.set_legacyPairing(objects["LegacyPairing"].Value().(bool))
                }

                if objects["Paired"].Value() != nil {
                    rd.set_paired(objects["Paired"].Value().(bool))
                }

                if objects["ServicesResolved"].Value() != nil {
                    rd.set_servicesResolved(objects["ServicesResolved"].Value().(bool))
                }

                if objects["Trusted"].Value() != nil {
                    rd.set_trusted(objects["Trusted"].Value().(bool))
                }

                if objects["ServiceData"].Value() != nil {
                    a := objects["ServiceData"].Value().(map[string]dbus.Variant)
                    b := make(map[string][]uint8)
                    for key, value := range a {
                        b[key]=value.Value().([]uint8)
                    }
                    rd.set_serviceData(b)
                }

                if objects["ManufacturerData"].Value() != nil {
                    a := objects["ManufacturerData"].Value().(map[uint16]dbus.Variant)
                    b := make(map[uint16][]uint8)
                    for key, value := range a {
                        b[key]=value.Value().([]uint8)
                    }
                    rd.set_manufacturerData(b)
                }

                if objects["RSSI"].Value() != nil {
                    rd.set_rssi(objects["RSSI"].Value().(int16))
                }

                if objects["TxPower"].Value() != nil {
                    rd.set_txpower(objects["TxPower"].Value().(int16))
                }

                if objects["Adapter"].Value() != nil {
                    rd.set_adapter(objects["Adapter"].Value().(dbus.ObjectPath))
                }

                if objects["Icon"].Value() != nil {
                    rd.set_icon(objects["Icon"].Value().(string))
                }

                if objects["Appearance"].Value() != nil {
                    rd.set_appearance(objects["Appearance"].Value().(uint16))
                }

                s.res[index] = *rd
                s.last = *rd

            case "org.freedesktop.DBus.ObjectManager.InterfacesRemoved":
                delete(s.res,string(v.Body[0].(dbus.ObjectPath)))
                //fmt.Printf("interface was removed \n")
            }
            //fmt.Printf("Size of Devices: %v \n", len(s.res))

            s.f()
        }
    //fmt.Printf("Size of Devices: %v \n", len(s.res))
    //fmt.Print("%v",s.res)
	}

    return nil
}

func (s *Scanner) Get_advertisements() map[string]map[string]interface{} {
    results := make(map[string]map[string]interface{})
    for _, value := range s.res {
        results[value.Get_address()] = make(map[string]interface{})
        results[value.Get_address()]["ManufacturerData"] = value.Get_manufacturerData()
        results[value.Get_address()]["Name"] = value.Get_name()
        //fmt.Printf("Address: %v\n",value.Get_address())
        //fmt.Printf("Name: %v\n",value.Get_name())
        //fmt.Printf("key type: %T \nvalue value: %v, type: %T \n \n",results,value.Get_manufacturerData(),value.Get_manufacturerData())
    }
    return results
}

func (s *Scanner) Get_last_advertisements() map[string]interface{} {
    results := make(map[string]interface{})
        results["Address"] = s.last.Get_address()
        results["ManufacturerData"] = s.last.Get_manufacturerData()
        results["Name"] = s.last.Get_name()
        //fmt.Printf("Address: %v\n",value.Get_address())
        //fmt.Printf("Name: %v\n",value.Get_name())
        //fmt.Printf("key type: %T \nvalue value: %v, type: %T \n \n",results,value.Get_manufacturerData(),value.Get_manufacturerData())

    return results
}
