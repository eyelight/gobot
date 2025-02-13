// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/drivers/gpio"
	"github.com/eyelight/gobot/platforms/dexter/gopigo3"
	"github.com/eyelight/gobot/platforms/raspi"
)

func main() {
	raspiAdaptor := raspi.NewAdaptor()
	gpg3 := gopigo3.NewDriver(raspiAdaptor)
	servo := gpio.NewServoDriver(gpg3, "SERVO_1")

	work := func() {
		gobot.Every(1*time.Second, func() {
			i := uint8(gobot.Rand(180))
			fmt.Println("Turning", i)
			servo.Move(i)
		})

	}

	robot := gobot.NewRobot("gopigo3servo",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gpg3, servo},
		work,
	)

	robot.Start()
}
