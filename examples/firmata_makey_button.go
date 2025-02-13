// +build example
//
// Do not build by default.

/*
 How to run
 Pass serial port to use as the first param:

	go run examples/firmata_makey_button.go /dev/ttyACM0
*/

package main

import (
	"os"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/drivers/gpio"
	"github.com/eyelight/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	button := gpio.NewMakeyButtonDriver(firmataAdaptor, "2")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			led.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("makeyBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{button, led},
		work,
	)

	robot.Start()
}
