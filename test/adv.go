package main

import (
	//"encoding/json"
	"github.com/godbus/dbus"
	//"github.com/godbus/dbus/prop"
	//"github.com/godbus/dbus/introspect"
	//"os"
	"fmt"
)

var BLUEZ_SERVICE_NAME = "org.bluez"
var LE_ADVERTISING_MANAGER_IFACE = "org.bluez.LEAdvertisingManager1"
var DBUS_OM_IFACE = "org.freedesktop.DBus.ObjectManager"
var DBUS_PROP_IFACE = "org.freedesktop.DBus.Properties"

var LE_ADVERTISEMENT_IFACE = "org.bluez.LEAdvertisement1"

type advertisement_package struct {
    ad_path string
    ad_bus string
    ad_type string
    ad_serviceUUIDs []string
    ad_manufacturerData map[uint16][]uint16
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

func (adv *advertisement_package) add_manufacturerData(manuf_code uint16, data []uint16) {
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

	//get adapter
	adapter_path := dbus.ObjectPath("/org/bluez/hci0")
	adapter := bus.Object(BLUEZ_SERVICE_NAME,adapter_path)
	fmt.Println(adapter.Path())
	fmt.Println(adapter.Destination())
	address,err := adapter.GetProperty("org.bluez.Adapter1.Alias")
	fmt.Printf("value: %v, err: %v \n",address.String(),err)
	fmt.Printf("value: %T, err: %v \n \n",address,err)

	var result interface{}
	powered1 := true
	err = adapter.Call("org.freedesktop.DBus.Properties.Set", 0, "org.bluez.Adapter1", "Powered", dbus.MakeVariant(powered1)).Store(&result)
	fmt.Printf("result value: %v, type: %T \n",result,result)
	fmt.Printf("error  value: %v, type: %T \n",err,err)

	/*props := map[string]map[string]*prop.Prop{
		"org.bluez.Adapter1": {
			"SomeInt": {
				int32(0),
				true,
				prop.EmitTrue,
				func(c *prop.Change) *dbus.Error {
					fmt.Println(c.Name, "changed to", c.Value)
					return nil
				},
			},
		},
	}*/


	//fmt.Println(props)

	//get adapter Properties
	//props := prop.New(bus,bus.Object(BLUEZ_SERVICE_NAME,adapter),props)


	adv := new(advertisement_package)

	adv.ad_type = "broadcast"
	adv.ad_serviceUUIDs = []string{"0x1800"}
	adv.ad_manufacturerData = map[uint16][]uint16{0x026B:{0xFFFF}}
	adv.ad_solicitUUIDs = make([]string,0)
	adv.ad_serviceData = map[string][]uint8{"0x1800":{0x01}}
	adv.ad_includeTxPower = false


}
