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
	a := api.NewAPI(master)
	a.Port = "8080"
	a.Start()

	pebbleAdaptor := pebble.NewAdaptor()
	pebbleDriver := pebble.NewDriver(pebbleAdaptor)

	work := func() {
		pebbleDriver.On(pebbleDriver.Event("accel"), func(data interface{}) {
			fmt.Println(data.(string))
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
