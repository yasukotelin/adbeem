package cmd

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/yasukotelin/adbeem/adb"
)

func Screencap(cli *cli.Context) error {
	adb := adb.New()

	if !adb.ExistsCommand() {
		return errors.New("adb command not found")
	}

	device, err := selectDeviceWithPrompt(adb)
	if err != nil {
		return err
	}

	adb.Device = device

	path, err := adb.Screencap()
	if err != nil {
		return err
	}

	output := cli.String("output")
	if err := adb.Pull(path, output); err != nil {
		return err
	}

	fmt.Println("screencap is success.")

	return nil
}
