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
	ollieBot := ollie.NewDriver(bleAdaptor)

	work := func() {
		head := 90
		ollieBot.SetRGB(255, 0, 0)
		ollieBot.Boost(true)
		gobot.Every(1*time.Second, func() {
			ollieBot.Roll(0, uint16(head))
			time.Sleep(1 * time.Second)
			head += 90
			head = head % 360
		})
	}

	robot := gobot.NewRobot("ollieBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ollieBot},
		work,
	)

	robot.Start()
}
