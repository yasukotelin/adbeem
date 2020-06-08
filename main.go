package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yasukotelin/adbeem/cmd"
)

func main() {
	app := &cli.App{
		Name:    "adbeem",
		Usage:   "adbeem is a CLI tool to capture or record the screen of Android with adb",
		Version: "0.1.0",
		Commands: []*cli.Command{
			{
				Name:   "screencap",
				Usage:  "capture the screen",
				Action: cmd.Screencap,
			},
			{
				Name:   "screenrecord",
				Usage:  "records the screen",
				Action: cmd.Screenrecord,
			},
		},
		Action: cli.ShowAppHelp,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
