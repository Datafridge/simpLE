package main

import (
  "github.com/esiqveland/notify"
  "github.com/mmbuw/simpLE"
  "github.com/godbus/dbus"
  "fmt"
)

type appconf struct {
    name string
    icon string
}

var scanner simpLE.Scanner

func main()  {
  scanner.Start(test_callback)
}

func test_callback() {
    results := scanner.Get_last_advertisements()

    if _, ok := results["ManufacturerData"].(map[uint16][]uint8)[65535]; ok {
        conn, err := dbus.SessionBus()
    	if err != nil {
    		panic(err)
        }

        manData := string(results["ManufacturerData"].(map[uint16][]uint8)[65535])
        manData1 := results["ManufacturerData"].(map[uint16][]uint8)[65535]
        fmt.Printf("ManData1: %v\n",manData1)

        if len(manData) == 21 {
            manData = manData + "..."
        }

        apps := map[uint8]appconf{
            1 : appconf{"GMail","mail-unread"},
            2 : appconf{"Snapchat","mail-unread"},
        }

        ticker := manData[3:]
        appnumber := results["ManufacturerData"].(map[uint16][]uint8)[65535][2]
        fmt.Printf("%v,%T\n",appnumber,appnumber)
        var appname string = "default"
        var iconName string = "mail-unread"


        if val, ok := apps[appnumber]; ok {
            appname = val.name
            iconName = val.icon
        }
        fmt.Printf("%v\n",appname)

        fmt.Printf("Ticker: %v\n",ticker)
        fmt.Printf("appname: %v\n",appname)

    	n := notify.Notification{
    		AppName:       "NotiLE",
    		ReplacesID:    uint32(0),
    		AppIcon:       iconName,
    		Summary:       appname,
    		Body:          ticker,
    		Actions:       []string{"cancel", "Cancel", "open", "Open"},
    		Hints:         map[string]dbus.Variant{},
    		ExpireTimeout: int32(5000),
    	}

        notify.SendNotification(conn, n)
    }

}
