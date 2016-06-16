package main

import (
	//"encoding/json"
	//"github.com/godbus/dbus"
	//"github.com/godbus/dbus/introspect"
	//"os"
	//"fmt"
)

type advertisement_package struct {
    ad_type string
    ad_serviceUUIDs []string
    manufacturerData map[string]string
    solicitUUIDs []string
    serviceData map[string]string
    includeTxPower bool
}

func main() {
	adv := new(advertisement_package)

	adv.ad_type = "broadcast"
	adv.ad_serviceUUIDs = []string{"0x1800"}
	adv.manufacturerData = map[string]string{"​0x026B":"node"}
  adv.solicitUUIDs = make([]string,0)
  adv.serviceData = map[string]string{"0x1800":"test"}
  adv.includeTxPower = false


}
