// +build example
//
// Do not build by default.

package main

import (
	"time"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/drivers/gpio"
	"github.com/eyelight/gobot/platforms/intel-iot/edison"
)

func main() {
	e := edison.NewAdaptor()
	led := gpio.NewLedDriver(e, "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{e},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
