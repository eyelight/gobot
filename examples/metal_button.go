// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/eyelight/gobot/drivers/gpio"
	"github.com/eyelight/gobot/platforms/intel-iot/edison"
)

func main() {
	e := edison.NewAdaptor()
	e.Connect()

	led := gpio.NewLedDriver(e, "13")
	led.Start()
	led.Off()

	button := gpio.NewButtonDriver(e, "5")
	button.Start()

	buttonEvents := button.Subscribe()
	for {
		select {
		case event := <-buttonEvents:
			fmt.Println("Event:", event.Name, event.Data)
			if event.Name == gpio.ButtonPush {
				led.Toggle()
			}
		}
	}
}
