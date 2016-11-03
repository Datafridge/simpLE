package simpLE

import (
    "github.com/godbus/dbus"
)

type Sender struct {
    dev *device
    adv *advertisement
    active bool
}

func (s *Sender) device(string error) {

}

func (s *Sender) Start(data string) error {
    var dev device
    var adv advertisement

    adv.ad_type = "peripheral"
	adv.ad_manufacturerData = map[uint16][]uint8{0xFFFF:[]byte(data)}

    adv.register(dev.dbusobject(),"advertisement1")

    var dic map[string]dbus.Variant
	err := dev.dbusobject().Call("org.bluez.LEAdvertisingManager1.RegisterAdvertisement", 0, adv.ad_path, dic).Store(&result)
    if err != nil {
        s.active = true
    }

    for s.active == true{}

    return nil
}
