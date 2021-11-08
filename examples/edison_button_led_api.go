// +build example
//
// Do not build by default.

package main

import (
	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/drivers/gpio"
	"github.com/eyelight/gobot/platforms/intel-iot/edison"

	"github.com/eyelight/gobot/api"
)

func main() {
	master := gobot.NewMaster()
	api.NewAPI(master).Start()

	e := edison.NewAdaptor()

	button := gpio.NewButtonDriver(e, "2")
	led := gpio.NewLedDriver(e, "4")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			led.On()
		})
		button.On(gpio.ButtonRelease, func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{e},
		[]gobot.Device{led, button},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
