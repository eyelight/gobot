package main

import (
	"os"

	"github.com/eyelight/gobot"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gobot"
	app.Author = "The Gobot team"
	app.Email = "https://github.com/eyelight/gobot"
	app.Version = gobot.Version()
	app.Usage = "Command Line Utility for generating new Gobot adaptors, drivers, and platforms"
	app.Commands = []cli.Command{
		Generate(),
	}
	app.Run(os.Args)
}
