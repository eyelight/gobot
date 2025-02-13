// +build example
//
// Do not build by default.

/*
 How to run
 Pass the Bluetooth address or name as the first param:

	go run examples/bb8.go BB-1234

 NOTE: sudo is required to use BLE in Linux
*/

package main

import (
	"os"
	"time"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/platforms/ble"
	"github.com/eyelight/gobot/platforms/sphero/bb8"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	bb8 := bb8.NewDriver(bleAdaptor)

	work := func() {
		gobot.Every(1*time.Second, func() {
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			bb8.SetRGB(r, g, b)
		})
	}

	robot := gobot.NewRobot("bbBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{bb8},
		work,
	)

	robot.Start()
}
