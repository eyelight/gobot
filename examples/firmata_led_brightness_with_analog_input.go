// +build example
//
// Do not build by default.

/*
 How to run
 Pass serial port to use as the first param:

	go run examples/firmata_led_brightness_with_analog_input.go /dev/ttyACM0
*/

package main

import (
	"fmt"
	"os"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/drivers/aio"
	"github.com/eyelight/gobot/drivers/gpio"
	"github.com/eyelight/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	sensor := aio.NewAnalogSensorDriver(firmataAdaptor, "0")
	led := gpio.NewLedDriver(firmataAdaptor, "3")

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			brightness := uint8(
				gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 1024), 0, 255),
			)
			fmt.Println("sensor", data)
			fmt.Println("brightness", brightness)
			led.Brightness(brightness)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor, led},
		work,
	)

	robot.Start()
}
