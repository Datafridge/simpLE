package simpLE

import (
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

func (d *device) set_path(i dbus.ObjectPath) {
    d.path = i
}

func (d *device) get_path() dbus.ObjectPath {
    return d.path
}


func (d *device) set_uuids(i []string) {
    d.uuids = i
}

func (d *device) get_uuids() []string {
    return d.uuids
}


func (d *device) set_address(i string) {
    d.address = i
}

func (d *device) get_address() string {
    return d.address
}


func (d *device) set_alias(i string) {
    d.alias = i
}

func (d *device) get_alias() string {
    return d.alias
}


func (d *device) set_modalias(i string) {
    d.modalias = i
}

func (d *device) get_modalias() string {
    return d.modalias
}


func (d *device) set_name(i string) {
    d.name = i
}

func (d *device) get_name() string {
    return d.name
}


func (d *device) set_class(i uint32) {
    d.class = i
}

func (d *device) get_class() uint32 {
    return d.class
}


func (d *local_device) set_le(i bool) {
    d.le = i
}

func (d *local_device) get_le() bool {
    return d.le
}


func (d *local_device) set_discoverable(i bool) {
    d.discoverable = i
}

func (d *local_device) get_discoverable() bool {
    return d.discoverable
}


func (d *local_device) set_discovering(i bool) {
    d.discovering = i
}

func (d *local_device) get_discovering() bool {
    return d.discovering
}


func (d *local_device) set_pairable(i bool) {
    d.pairable = i
}

func (d *local_device) get_pairable() bool {
    return d.pairable
}


func (d *local_device) set_powered(i bool) {
    d.powered = i
}

func (d *local_device) get_powered() bool {
    return d.powered
}


func (d *local_device) set_discoverableTimeout(i uint32) {
    d.discoverableTimeout = i
}

func (d *local_device) get_discoverableTimeout() uint32 {
    return d.discoverableTimeout
}


func (d *local_device) set_pairableTimeout(i uint32) {
    d.pairableTimeout = i
}

func (d *local_device) get_pairableTimeout() uint32 {
    return d.pairableTimeout
}


func (d *remote_device) set_blocked(i bool) {
    d.blocked = i
}

func (d *remote_device) get_blocked() bool {
    return d.blocked
}


func (d *remote_device) set_connected(i bool) {
    d.connected = i
}

func (d *remote_device) get_connected() bool {
    return d.connected
}


func (d *remote_device) set_legacyPairing(i bool) {
    d.legacyPairing = i
}

func (d *remote_device) get_legacyPairing() bool {
    return d.legacyPairing
}


func (d *remote_device) set_paired(i bool) {
    d.paired = i
}

func (d *remote_device) get_paired() bool {
    return d.paired
}


func (d *remote_device) set_servicesResolved(i bool) {
    d.servicesResolved = i
}

func (d *remote_device) get_servicesResolved() bool {
    return d.servicesResolved
}


func (d *remote_device) set_trusted(i bool) {
    d.trusted = i
}

func (d *remote_device) get_trusted() bool {
    return d.trusted
}


func (d *remote_device) set_serviceData(i map[string][]uint8) {
    d.serviceData = i
}

func (d *remote_device) get_serviceData() map[string][]uint8 {
    return d.serviceData
}


func (d *remote_device) set_manufacturerData(i map[uint16][]uint8) {
    d.manufacturerData = i
}

func (d *remote_device) get_manufacturerData() map[uint16][]uint8 {
    return d.manufacturerData
}


func (d *remote_device) set_rssi(i int16) {
    d.rssi = i
}

func (d *remote_device) get_rssi() int16 {
    return d.rssi
}


func (d *remote_device) set_txpower(i int16) {
    d.txpower = i
}

func (d *remote_device) get_txpower() int16 {
    return d.txpower
}


func (d *remote_device) set_adapter(i dbus.ObjectPath) {
    d.adapter = i
}

func (d *remote_device) get_adapter() dbus.ObjectPath {
    return d.adapter
}


func (d *remote_device) set_icon(i string) {
    d.icon = i
}

func (d *remote_device) get_icon() string {
    return d.icon
}


func (d *remote_device) set_appearance(i uint16) {
    d.appearance = i
}

func (d *remote_device) get_appearance() uint16 {
    return d.appearance
}


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
