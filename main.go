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
		Usage:   "adbeem is a CLI tool using Android adb command more easily.",
		Version: "1.6.0",
		Commands: []*cli.Command{
			{
				Name:    "screencap",
				Aliases: []string{"screenshot"},
				Usage:   "capture the screen",
				Action:  cmd.Screencap,
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
						Usage:   "include converted gif",
					},
					&cli.StringFlag{
						Name:  "gif-quality",
						Usage: "Gif quolity. middle or high.",
						Value: "middle",
					},
					&cli.StringFlag{
						Name:  "gif-orientation",
						Usage: "Gif orientation. portlait or landscape",
						Value: "portlait",
					},
					&cli.StringFlag{
						Name:  "fps",
						Value: "20",
						Usage: "Gif fps",
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
						Usage:    "deep link url. (ex. \"example://app\")",
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
