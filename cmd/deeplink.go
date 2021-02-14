package cmd

import (
	"errors"

	"github.com/urfave/cli/v2"
	"github.com/yasukotelin/adbeem/adb"
)

func DeepLink(cli *cli.Context) error {
	adb := adb.New()

	if !adb.ExistsCommand() {
		return errors.New("adb command not found")
	}

	device, err := selectDeviceWithPrompt(adb)
	if err != nil {
		return err
	}

	adb.Device = device

	url := cli.String("url")

	if err = adb.SendDeepLink(url); err != nil {
		return err
	}

	return nil
}
