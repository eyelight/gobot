// +build example
//
// Do not build by default.

/*
 How to run
 Pass serial port to use as the first param:

	go run examples/firmata_mpl115a2.go /dev/ttyACM0
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
	mpl115a2 := i2c.NewMPL115A2Driver(firmataAdaptor)

	work := func() {
		gobot.Every(1*time.Second, func() {
			press, _ := mpl115a2.Pressure()
			fmt.Println("Pressure", press)

			temp, _ := mpl115a2.Temperature()
			fmt.Println("Temperature", temp)
		})
	}

	robot := gobot.NewRobot("mpl115a2Bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{mpl115a2},
		work,
	)

	robot.Start()
}
