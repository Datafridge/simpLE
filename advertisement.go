package simpLE

import ()

type advertisement struct {
    ad_path string
    ad_bus string
    ad_type string
    ad_serviceUUIDs []string
    ad_manufacturerData map[uint16][]uint8
    ad_solicitUUIDs []string
    ad_serviceData map[string][]uint8
    ad_includeTxPower bool
}

func (adv *advertisement) add_serviceUUIDs(uuid string) {
    adv.ad_serviceUUIDs = append(adv.ad_serviceUUIDs,uuid)
}

func (adv *advertisement) add_solicitUUIDs(uuid string) {
    adv.ad_solicitUUIDs = append(adv.ad_solicitUUIDs,uuid)
}

func (adv *advertisement) add_manufacturerData(manuf_code uint16, data []uint8) {
    adv.ad_manufacturerData[manuf_code] = data
}

func (adv *advertisement) add_serviceData(uuid string, data []uint8) {
    adv.ad_serviceData[uuid] = data
}
