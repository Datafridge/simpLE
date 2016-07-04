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
    ad_path string
    ad_bus string
    ad_type string
    ad_serviceUUIDs []string
    ad_manufacturerData map[uint8][]uint8
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

func (adv *advertisement_package) add_manufacturerData(manuf_code uint8, data []uint8) {
    adv.ad_manufacturerData[manuf_code] = data
}

func (adv *advertisement_package) add_serviceData(uuid string, data []uint8) {
    adv.ad_serviceData[uuid] = data
}

func main() {

}
