// +build example
//
// Do not build by default.

package main

import (
	"time"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/api"
	"github.com/eyelight/gobot/drivers/gpio"
	"github.com/eyelight/gobot/platforms/firmata"
)

func main() {
	master := gobot.NewMaster()
	a := api.NewAPI(master)
	a.Start()

	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
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

	master.AddRobot(robot)

	master.Start()
}
