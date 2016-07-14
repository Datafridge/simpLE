package main

import (
	//"encoding/json"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/prop"
	"github.com/godbus/dbus/introspect"
	//"os"
	"fmt"
)

var BLUEZ_SERVICE_NAME = "org.bluez"
var LE_ADVERTISING_MANAGER_IFACE = "org.bluez.LEAdvertisingManager1"
var DBUS_OM_IFACE = "org.freedesktop.DBus.ObjectManager"
var DBUS_PROP_IFACE = "org.freedesktop.DBus.Properties"

var LE_ADVERTISEMENT_IFACE = "org.bluez.LEAdvertisement1"

type foo string

func (f foo) Foo() (string, *dbus.Error) {
	fmt.Println(f)
	return string(f), nil
}

type advertisement_package struct {
    ad_path string
    ad_bus string
    ad_type string
    ad_serviceUUIDs []string
    ad_manufacturerData map[uint16][]uint8
    ad_solicitUUIDs []string
    ad_serviceData map[string][]uint8
    ad_includeTxPower bool
    // DBUS Service Object
}

func (adv *advertisement_package) add_serviceUUIDs(uuid string) {
    adv.ad_serviceUUIDs = append(adv.ad_serviceUUIDs,uuid)
}

func (adv *advertisement_package) add_solicitUUIDs(uuid string) {
    adv.ad_solicitUUIDs = append(adv.ad_solicitUUIDs,uuid)
}

func (adv *advertisement_package) add_manufacturerData(manuf_code uint16, data []uint8) {
    adv.ad_manufacturerData[manuf_code] = data
}

func (adv *advertisement_package) add_serviceData(uuid string, data []uint8) {
    adv.ad_serviceData[uuid] = data
}

func main() {
	//connect to systembus
	bus, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}

	/*sbus, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}*/

	//get adapter
	adapter_path := dbus.ObjectPath("/org/bluez/hci0")
	adapter := bus.Object(BLUEZ_SERVICE_NAME,adapter_path)
	//fmt.Println(adapter.Path())
	//fmt.Println(adapter.Destination())
	//address,err := adapter.GetProperty("org.bluez.Adapter1.Alias")
	//fmt.Printf("value: %v, err: %v \n",address.String(),err)
	//fmt.Printf("value: %T, err: %v \n \n",address,err)

	var result interface{}
	powered1 := true
	err = adapter.Call("org.freedesktop.DBus.Properties.Set", 0, "org.bluez.Adapter1", "Powered", dbus.MakeVariant(powered1)).Store(&result)
	//fmt.Printf("result value: %v, type: %T \n",result,result)
	//fmt.Printf("error  value: %v, type: %T \n",err,err)

	//adapter_manager_path := dbus.ObjectPath("/org/bluez/hci0")
	//adapter := bus.Object(BLUEZ_SERVICE_NAME,adapter_manager_path)

	adv := new(advertisement_package)

	adv.ad_type = "peripheral"
	adv.ad_serviceUUIDs = []string{"180D","180F"}
	adv.ad_manufacturerData = map[uint16][]uint8{0xFFFF:{0x00, 0x01, 0x02, 0x03, 0x04}}
	//adv.ad_solicitUUIDs = make([]string,0)
	adv.ad_serviceData = map[string][]uint8{"9999":{0x00, 0x01, 0x02, 0x03, 0x04}}
	adv.ad_includeTxPower = true

	/*adv.ad_type = "peripheral"
	adv.ad_serviceUUIDs = []string{"180D","180F"}
	adv.ad_manufacturerData = map[uint16][]uint8{0xFFFF:{0x00, 0x01, 0x02, 0x03, 0x04}}
	//adv.ad_solicitUUIDs = make([]string,0)
	adv.ad_serviceData = map[string][]uint8{"9999":{0x00, 0x01, 0x02, 0x03, 0x04}}
	adv.ad_includeTxPower = true*/

	//TODO: request name before add new object

	advertisement_path := dbus.ObjectPath("/org/bluez/simpLE/advertisement1")
	//advertisement1 := bus.Object(BLUEZ_SERVICE_NAME,advertisement_path)
	f := foo("Bar")
	bus.Export(f, advertisement_path, "org.bluez.LEAdvertisement1")


	//add property
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

	//fmt.Println("org.bluez.LEAdvertisement1 / /org/bluez/simpLE/advertisement1 ...")

	/*c := make(chan *dbus.Signal)
	bus.Signal(c)
	for _ = range c {
	}*/

	var dic map[string]dbus.Variant
	err = adapter.Call("org.bluez.LEAdvertisingManager1.RegisterAdvertisement", 0, advertisement_path, dic).Store(&result)
	//fmt.Printf("result: value: %T, err: %v \n \n",result,result)
	fmt.Printf("err: value: %T, err: %v \n \n",err,err)

	//value1,err := advertisement1.GetProperty("org.bluez.LEAdvertisement1.Type")
	//fmt.Printf("value1: value: %T, err: %v \n \n",value1.String(),value1.String())


}
