package simpLE

import (
	//"encoding/json"
	//"github.com/godbus/dbus"
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
    manufacturerData map[uint8]uint8
    solicitUUIDs []string
    serviceData map[uint8]uint8
    includeTxPower bool
}

func main() {

}
