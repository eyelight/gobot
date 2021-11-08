/*
Package beaglebone provides the Gobot adaptor for the Beaglebone Black/Green, as well as a
separate Adaptor for the PocketBeagle.

Installing:

	go get github.com/eyelight/gobot/platforms/beaglebone

Example:

	package main

	import (
		"time"

		"github.com/eyelight/gobot"
		"github.com/eyelight/gobot/drivers/gpio"
		"github.com/eyelight/gobot/platforms/beaglebone"
	)

	func main() {
		beagleboneAdaptor := beaglebone.NewAdaptor()
		led := gpio.NewLedDriver(beagleboneAdaptor, "P9_12")

		work := func() {
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("blinkBot",
			[]gobot.Connection{beagleboneAdaptor},
			[]gobot.Device{led},
			work,
		)

		robot.Start()
	}

For more information refer to the beaglebone README:
https://github.com/hybridgroup/gobot/blob/master/platforms/beaglebone/README.md
*/
package beaglebone // import "github.com/eyelight/gobot/platforms/beaglebone"
