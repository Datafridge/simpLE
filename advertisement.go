package advertisement

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
    manufacturerData map[uint8]uint8
    solicitUUIDs []string
    serviceData map[uint8]uint8
    includeTxPower bool
}

func main() {

}
