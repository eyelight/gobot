// +build example
//
// Do not build by default.

/*
 How to run
 Pass the BLE address or BLE name as first param:

	go run examples/ble_firmata_blink.go FIRMATA

 NOTE: sudo is required to use BLE in Linux
*/

package main

import (
	"os"
	"time"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/drivers/gpio"
	"github.com/eyelight/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewBLEAdaptor(os.Args[1])
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
