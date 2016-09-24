package main

import (
  "github.com/esiqveland/notify"
  "github.com/mmbuw/simpLE"
  "github.com/godbus/dbus"
  "fmt"
)

var scanner simpLE.Scanner

func main()  {
  scanner.Start(test_callback)
}

func test_callback() {
    results := scanner.Get_last_advertisements()

    //fmt.Printf("Address: %v\n",results["Address"])
    //fmt.Printf("Name: %v\n",results["Name"])
    //fmt.Printf("ManufacturerData: %v \n\n",string(results["ManufacturerData"].(map[uint16][]uint8)[65535]))

    if _, ok := results["ManufacturerData"].(map[uint16][]uint8)[65535]; ok {
        conn, err := dbus.SessionBus()
    	if err != nil {
    		panic(err)
        }

        manData := string(results["ManufacturerData"].(map[uint16][]uint8)[65535])

        if len(manData) == 21 {
            manData = manData + "..."
        }

        ticker := manData[3:]
        appnumber := results["ManufacturerData"].(map[uint16][]uint8)[65535][2]
        fmt.Printf("%v,%T",appnumber,appnumber)
        var appname string


        switch appnumber {
        case 1 :
            appname = "GMail"
        }

    	iconName := "mail-unread"
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
