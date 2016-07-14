package main

import (
	"fmt"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
	"github.com/godbus/dbus/prop"
	"os"
)

type release string

func (r release) Release() (*dbus.Error) {
	fmt.Println("cleaned")
	return nil
}

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	reply, err := conn.RequestName("org.bluez.LEAdvertisement1",
		dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		fmt.Fprintln(os.Stderr, "name already taken")
		os.Exit(1)
	}
    propsSpec := map[string]map[string]*prop.Prop{
		"org.bluez.LEAdvertisement1": {
			"Type": {
				string("peripheral"),
				true,
				prop.EmitTrue,
				func(c *prop.Change) *dbus.Error {
					fmt.Println(c.Name, "changed to", c.Value)
					return nil
				},
			},
			"ServiceUUIDs": {
				[]string([]string{"180D","180F"}),
				true,
				prop.EmitTrue,
				func(c *prop.Change) *dbus.Error {
					fmt.Println(c.Name, "changed to", c.Value)
					return nil
				},
			},
			"ManufacturerData": {
				map[uint16][]uint8(map[uint16][]uint8{0xFFFF:{0x00, 0x01, 0x02, 0x03, 0x04}}),
				true,
				prop.EmitTrue,
				func(c *prop.Change) *dbus.Error {
					fmt.Println(c.Name, "changed to", c.Value)
					return nil
				},
			},
			"SolicitUUIDs": {
				[]string(nil),
				true,
				prop.EmitTrue,
				func(c *prop.Change) *dbus.Error {
					fmt.Println(c.Name, "changed to", c.Value)
					return nil
				},
			},
			"ServiceData": {
				map[string][]uint8(map[string][]uint8{"9999":{0x00, 0x01, 0x02, 0x03, 0x04}}),
				true,
				prop.EmitTrue,
				func(c *prop.Change) *dbus.Error {
					fmt.Println(c.Name, "changed to", c.Value)
					return nil
				},
			},
			"IncludeTxPower": {
				bool(true),
				true,
				prop.EmitTrue,
				func(c *prop.Change) *dbus.Error {
					fmt.Println(c.Name, "changed to", c.Value)
					return nil
				},
			},
		},
	}
	r := release("i")
	conn.Export(r, "/org/bluez/simpLE/advertisement1", "org.bluez.LEAdvertisement1")
	props := prop.New(conn, "/org/bluez/simpLE/advertisement1", propsSpec)
	n := &introspect.Node{
		Name: "/org/bluez/simpLE/advertisement1",
		Interfaces: []introspect.Interface{
			introspect.IntrospectData,
			prop.IntrospectData,
			{
				Name:       "org.bluez.LEAdvertisement1",
				Methods:    introspect.Methods(r),
				Properties: props.Introspection("org.bluez.LEAdvertisement1"),
			},
		},
	}
	conn.Export(introspect.NewIntrospectable(n), "/org/bluez/simpLE/advertisement1",
		"org.freedesktop.DBus.Introspectable")
	fmt.Println("Listening on org.bluez.LEAdvertisement1 / /org/bluez/simpLE/advertisement1 ...")

	c := make(chan *dbus.Signal)
	conn.Signal(c)
	for _ = range c {
	}
}
