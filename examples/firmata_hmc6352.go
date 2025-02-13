// +build example
//
// Do not build by default.

/*
 How to run
 Pass serial port to use as the first param:

	go run examples/firmata_hmc6352.go /dev/ttyACM0
*/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/drivers/i2c"
	"github.com/eyelight/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	hmc6352 := i2c.NewHMC6352Driver(firmataAdaptor)

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			heading, _ := hmc6352.Heading()
			fmt.Println("Heading", heading)
		})
	}

	robot := gobot.NewRobot("hmc6352Bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{hmc6352},
		work,
	)

	robot.Start()
}
