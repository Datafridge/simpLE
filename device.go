package simpLE

import (
    //"fmt"
    "encoding/json"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
)

type device struct {
    path     dbus.ObjectPath
    uuids    []string
    address  string
    alias    string
    modalias string
    name     string
    class    uint32
}

type local_device struct {
    device
    le                  bool
    discoverable        bool
    discovering         bool
    pairable            bool
    powered             bool
    discoverableTimeout uint32
    pairableTimeout     uint32
}

type remote_device struct {
    device
    blocked          bool
    connected        bool
    legacyPairing    bool
    paired           bool
    servicesResolved bool
    trusted          bool
    serviceData      map[string][]uint8
    manufacturerData map[uint16][]uint8
    rssi             int16
    txpower          int16
    adapter          dbus.ObjectPath
    icon             string
    appearance       uint16
}

// Path - GET
func (d *device) get_path() dbus.ObjectPath {
    return d.path
}


// UUIDs - GET
func (d *device) get_uuids() []string {
    return d.uuids
}


// Address - GET
func (d *device) Get_address() string {
    return d.address
}


// Alias - SET/GET
func (d *device) set_alias(i string) {
    d.alias = i
}

func (d *device) get_alias() string {
    return d.alias
}


// Modalias - GET
func (d *device) get_modalias() string {
    return d.modalias
}


// Name - GET
func (d *device) Get_name() string {
    return d.name
}


// Class - GET
func (d *device) get_class() uint32 {
    return d.class
}


// LE - GET
func (d *local_device) get_le() bool {
    return d.le
}


// Discoverable - SET/GET
func (d *local_device) set_discoverable(i bool) {
    d.discoverable = i
}

func (d *local_device) get_discoverable() bool {
    return d.discoverable
}


// Discovering - GET
func (d *local_device) get_discovering() bool {
    return d.discovering
}


// Pairable - SET/GET
func (d *local_device) set_pairable(i bool) {
    d.pairable = i
}

func (d *local_device) get_pairable() bool {
    return d.pairable
}


// Powered - SET/GET
func (d *local_device) set_powered(i bool) {
    d.powered = i
}

func (d *local_device) get_powered() bool {
    return d.powered
}


// DiscoverableTimeout - SET/GET
func (d *local_device) set_discoverableTimeout(i uint32) {
    d.discoverableTimeout = i
}

func (d *local_device) get_discoverableTimeout() uint32 {
    return d.discoverableTimeout
}

// PairableTimeout - SET/GET
func (d *local_device) set_pairableTimeout(i uint32) {
    d.pairableTimeout = i
}

func (d *local_device) get_pairableTimeout() uint32 {
    return d.pairableTimeout
}


// Blocked - SET/GET
func (d *remote_device) set_blocked(i bool) {
    d.blocked = i
}

func (d *remote_device) get_blocked() bool {
    return d.blocked
}


// Connected - SET/GET
func (d *remote_device) get_connected() bool {
    return d.connected
}


// LegacyPairing - GET
func (d *remote_device) get_legacyPairing() bool {
    return d.legacyPairing
}


// Paired - GET
func (d *remote_device) get_paired() bool {
    return d.paired
}


// ServicesResolved - GET
func (d *remote_device) get_servicesResolved() bool {
    return d.servicesResolved
}


// Trusted - SET/GET
func (d *remote_device) set_trusted(i bool) {
    d.trusted = i
}

func (d *remote_device) get_trusted() bool {
    return d.trusted
}


// ServiceData - GET
func (d *remote_device) get_serviceData() map[string][]uint8 {
    return d.serviceData
}


// ManufacturerData - GET
func (d *remote_device) Get_manufacturerData() map[uint16][]uint8 {
    return d.manufacturerData
}


// RSSI - GET
func (d *remote_device) get_rssi() int16 {
    return d.rssi
}


// TxPower - GET
func (d *remote_device) get_txpower() int16 {
    return d.txpower
}


// Adapter - SET/GET
func (d *remote_device) set_adapter(i dbus.ObjectPath) {
    d.adapter = i
}

func (d *remote_device) get_adapter() dbus.ObjectPath {
    return d.adapter
}


// Icon - GET
func (d *remote_device) get_icon() string {
    return d.icon
}


// Appearance - SET/GET
func (d *remote_device) get_appearance() uint16 {
    return d.appearance
}


// dbusobject - SET/GET
func (d *device) dbusobject()  dbus.BusObject {

    conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}

	adapter := conn.Object("org.bluez",dbus.ObjectPath("/org/bluez/hci0"))

    return adapter
}

type temp_Device struct {
    XMLName string
		Name string
		Interfaces interface {}
		Children []struct {
			XMLName interface {}
			Name string
		}
}

// get all available devices
//TODO auch auf LE prüfen
func Get_devices() string {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	node, err := introspect.Call(conn.Object("org.bluez", "/org/bluez"))
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(node, "", "    ")
	var device temp_Device

	json.Unmarshal(data, &device)
	return device.Children[0].Name
}
