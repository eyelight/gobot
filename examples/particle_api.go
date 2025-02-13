// +build example
//
// Do not build by default.

/*
 To run this example, pass the device ID as first param,
 and the access token as the second param:

	go run examples/particle_api.go mydevice myaccesstoken
*/

package main

import (
	"os"
	"time"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/api"
	"github.com/eyelight/gobot/drivers/gpio"
	"github.com/eyelight/gobot/platforms/particle"
)

func main() {
	master := gobot.NewMaster()
	api.NewAPI(master).Start()

	core := particle.NewAdaptor(os.Args[1], os.Args[2])
	led := gpio.NewLedDriver(core, "D7")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("spark",
		[]gobot.Connection{core},
		[]gobot.Device{led},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
