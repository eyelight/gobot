// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/api"
	"github.com/eyelight/gobot/platforms/pebble"
)

func main() {
	master := gobot.NewMaster()
	api := api.NewAPI(master)
	api.Port = "8080"
	api.Start()

	pebbleAdaptor := pebble.NewAdaptor()
	pebbleDriver := pebble.NewDriver(pebbleAdaptor)

	work := func() {
		pebbleDriver.SendNotification("Hello Pebble!")
		pebbleDriver.On(pebbleDriver.Event("button"), func(data interface{}) {
			fmt.Println("Button pushed: " + data.(string))
		})

		pebbleDriver.On(pebbleDriver.Event("tap"), func(data interface{}) {
			fmt.Println("Tap event detected")
		})
	}

	robot := gobot.NewRobot("pebble",
		[]gobot.Connection{pebbleAdaptor},
		[]gobot.Device{pebbleDriver},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
