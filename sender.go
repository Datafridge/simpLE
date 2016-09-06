package simpLE

import (
    "github.com/godbus/dbus"
    //"time"
    //"errors"
    //"fmt"
)

type Sender struct {
    dev *device
    adv *advertisement
    active bool
}

func (s *Sender) device(string error) {

}

func (s *Sender) Start() error {
    var dev device
    var adv advertisement

    adv.ad_type = "peripheral"
	adv.ad_serviceUUIDs = []string{"180D","180F"}
	//adv.ad_manufacturerData = map[uint16][]uint8{0xFFFF:{0xFF, 0x01, 0x02, 0x03, 0x04}}
	//adv.ad_solicitUUIDs = make([]string,0)
	adv.ad_serviceData = map[string][]uint8{"9999":{0x00, 0x01, 0x02, 0x03, 0x04}}
	adv.ad_includeTxPower = true

    adv.register(dev.dbusobject(),"advertisement1")

    var dic map[string]dbus.Variant
	err := dev.dbusobject().Call("org.bluez.LEAdvertisingManager1.RegisterAdvertisement", 0, adv.ad_path, dic).Store(&result)
    if err != nil {
        s.active = true
    }

    for s.active == true{}

    return nil
}
