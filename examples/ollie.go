// +build example
//
// Do not build by default.

package main

import (
	"os"
	"time"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/platforms/ble"
	"github.com/eyelight/gobot/platforms/sphero/ollie"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ollie := ollie.NewDriver(bleAdaptor)

	work := func() {
		gobot.Every(1*time.Second, func() {
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			ollie.SetRGB(r, g, b)
		})
	}

	robot := gobot.NewRobot("ollieBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ollie},
		work,
	)

	robot.Start()
}
