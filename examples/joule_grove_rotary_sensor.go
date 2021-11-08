// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/drivers/aio"
	"github.com/eyelight/gobot/drivers/i2c"
	"github.com/eyelight/gobot/platforms/intel-iot/joule"
)

func main() {
	board := joule.NewAdaptor()
	ads1015 := i2c.NewADS1015Driver(board)
	sensor := aio.NewGroveRotaryDriver(ads1015, "0")

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			fmt.Println("sensor", data)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{board},
		[]gobot.Device{ads1015, sensor},
		work,
	)

	robot.Start()
}
