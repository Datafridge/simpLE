package simpLE

import (
    "github.com/godbus/dbus"
    "github.com/godbus/dbus/prop"
    "github.com/godbus/dbus/introspect"
    "fmt"
)

var result interface{}

type advertisement struct {
    ad_path dbus.ObjectPath
    ad_bus string
    ad_type string
    ad_serviceUUIDs []string
    ad_manufacturerData map[uint16][]uint8
    ad_solicitUUIDs []string
    ad_serviceData map[string][]uint8
    ad_includeTxPower bool
}

type foo string

func (adv *advertisement) add_serviceUUIDs(uuid string) {
    adv.ad_serviceUUIDs = append(adv.ad_serviceUUIDs,uuid)
}

func (adv *advertisement) add_solicitUUIDs(uuid string) {
    adv.ad_solicitUUIDs = append(adv.ad_solicitUUIDs,uuid)
}

func (adv *advertisement) add_manufacturerData(manuf_code uint16, data []uint8) {
    adv.ad_manufacturerData[manuf_code] = data
}

func (adv *advertisement) add_serviceData(uuid string, data []uint8) {
    adv.ad_serviceData[uuid] = data
}

func (adv *advertisement) register(adapter dbus.BusObject, name string)  {

    bus, err := dbus.SystemBus()
    if err != nil {
        panic(err)
    }

    f := string("Bar")
    adv.ad_path = dbus.ObjectPath("/org/bluez/simpLE/"+name)

    propsSpec := map[string]map[string]*prop.Prop{
        "org.bluez.LEAdvertisement1": {
            "Type": {
                string(adv.ad_type),
                true,
                prop.EmitTrue,
                func(c *prop.Change) *dbus.Error {
                    fmt.Println(c.Name, "changed to", c.Value)
                    return nil
                },
            },
            "ServiceUUIDs": {
                []string(adv.ad_serviceUUIDs),
                true,
                prop.EmitTrue,
                func(c *prop.Change) *dbus.Error {
                    fmt.Println(c.Name, "changed to", c.Value)
                    return nil
                },
            },
            "ManufacturerData": {
                map[uint16][]uint8(adv.ad_manufacturerData),
                true,
                prop.EmitTrue,
                func(c *prop.Change) *dbus.Error {
                    fmt.Println(c.Name, "changed to", c.Value)
                    return nil
                },
            },
            "SolicitUUIDs": {
                []string(adv.ad_solicitUUIDs),
                true,
                prop.EmitTrue,
                func(c *prop.Change) *dbus.Error {
                    fmt.Println(c.Name, "changed to", c.Value)
                    return nil
                },
            },
            "ServiceData": {
                map[string][]uint8(adv.ad_serviceData),
                true,
                prop.EmitTrue,
                func(c *prop.Change) *dbus.Error {
                    fmt.Println(c.Name, "changed to", c.Value)
                    return nil
                },
            },
            "IncludeTxPower": {
                bool(adv.ad_includeTxPower),
                true,
                prop.EmitTrue,
                func(c *prop.Change) *dbus.Error {
                    fmt.Println(c.Name, "changed to", c.Value)
                    return nil
                },
            },
        },
    }

    props := prop.New(bus, "/org/bluez/simpLE/advertisement1", propsSpec)

    n := &introspect.Node{
		Name: "/org/bluez/simpLE/advertisement1",
		Interfaces: []introspect.Interface{
			introspect.IntrospectData,
			prop.IntrospectData,
			{
				Name:       "org.bluez.LEAdvertisement1",
				Methods:    introspect.Methods(f),
				Properties: props.Introspection("org.bluez.LEAdvertisement1"),
			},
		},
	}

    bus.Export(introspect.NewIntrospectable(n), "/org/bluez/simpLE/advertisement1","org.freedesktop.DBus.Introspectable")

    var dic map[string]dbus.Variant
	err = adapter.Call("org.bluez.LEAdvertisingManager1.RegisterAdvertisement", 0, adv.ad_path, dic).Store(&result)
}

func (f foo) Foo() (string, *dbus.Error) {
    fmt.Println(f)
    return string(f), nil
}
