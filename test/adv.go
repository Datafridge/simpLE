package main

import (
	//"encoding/json"
	"github.com/godbus/dbus"
	//"github.com/godbus/dbus/introspect"
	//"os"
	//"fmt"
)

var BLUEZ_SERVICE_NAME = "org.bluez"
var LE_ADVERTISING_MANAGER_IFACE = "org.bluez.LEAdvertisingManager1"
var DBUS_OM_IFACE = "org.freedesktop.DBus.ObjectManager"
var DBUS_PROP_IFACE = "org.freedesktop.DBus.Properties"

var LE_ADVERTISEMENT_IFACE = "org.bluez.LEAdvertisement1"

type advertisement_package struct {
    ad_type string
    ad_serviceUUIDs []string
    manufacturerData map[string]string
    solicitUUIDs []string
    serviceData map[string]string
    includeTxPower bool
}

func main() {
  bus, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}

  //device := new(dbus.ObjectPath)
  //device = *dbus.ObjectPath("/org/bluez/hci1")
  adapter := "/org/bluez/hci1"

  adv := new(advertisement_package)

	adv.ad_type = "broadcast"
	adv.ad_serviceUUIDs = []string{"0x1800"}
	adv.manufacturerData = map[string]string{"​0x026B":"node"}
  adv.solicitUUIDs = make([]string,0)
  adv.serviceData = map[string]string{"0x1800":"test"}
  adv.includeTxPower = false


}
