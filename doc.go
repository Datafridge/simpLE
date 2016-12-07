/*simpLE implements a easy to use Bluetooth Low-Energy API on Linux

Requirements:
dbus
bluez 5.40 or higher, with enabled experimental features

Usage:
Sender:
    package main

    import (
        "github.com/mmbuw/simpLE"
    )

    var sender simpLE.Sender

    func main()  {
        sender.Start("0123456789abcdef")
    }

Scanner:
    package main

    import (
        "github.com/mmbuw/simpLE"
        "fmt"
    )

    var scanner simpLE.Scanner
    var sender simpLE.Sender

    func main()  {
        scanner.Start(print_information)
        go sender.Start()
    }

    func print_information() {
        results := scanner.Get_last_advertisements()

        fmt.Printf("Address: %v\n",results["Address"])
        fmt.Printf("Name: %v\n",results["Name"])
        fmt.Printf("ManufacturerData: %v \n\n",results["ManufacturerData"].(map[uint16][]uint8)[65535])
    }

*/
package simpLE
