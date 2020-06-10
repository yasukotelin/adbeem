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
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:   "screencap",
				Usage:  "capture the screen",
				Action: cmd.Screencap,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "output path",
					},
				},
			},
			{
				Name:   "screenrecord",
				Usage:  "records the screen",
				Action: cmd.Screenrecord,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "output path",
					},
				},
			},
		},
		Action: cli.ShowAppHelp,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
