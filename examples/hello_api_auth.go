// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/api"
)

func main() {
	master := gobot.NewMaster()

	a := api.NewAPI(master)
	a.AddHandler(api.BasicAuth("gort", "klatuu"))
	a.Debug()

	a.AddHandler(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
	})
	a.Start()

	master.AddCommand("custom_gobot_command",
		func(params map[string]interface{}) interface{} {
			return "This command is attached to the mcp!"
		})

	hello := master.AddRobot(gobot.NewRobot("hello"))

	hello.AddCommand("hi_there", func(params map[string]interface{}) interface{} {
		return fmt.Sprintf("This command is attached to the robot %v", hello.Name)
	})

	master.Start()
}
