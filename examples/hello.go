// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/eyelight/gobot"
)

func main() {
	robot := gobot.NewRobot(
		func() {
			gobot.Every(500*time.Millisecond, func() { fmt.Println("Greetings human") })
		},
	)

	robot.Start()
}
