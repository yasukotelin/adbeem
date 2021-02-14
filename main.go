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
		Version: "1.2.0",
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
					&cli.BoolFlag{
						Name:    "gif",
						Aliases: []string{"g"},
						Usage:   "convert to gif",
					},
					&cli.StringFlag{
						Name:    "gifrate",
						Aliases: []string{"gr"},
						Value:   "15",
						Usage:   "Gif rate",
					},
				},
			},
			{
				Name:   "deeplink",
				Usage:  "send deeplink",
				Action: cmd.DeepLink,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "url",
						Aliases:  []string{"u"},
						Required: true,
						Usage:    "deep link url",
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
