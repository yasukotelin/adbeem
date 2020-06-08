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

	devices, err := adb.GetDevices()
	if err != nil {
		return err
	}

	// TODO devicesがEmptyの場合
	var device string
	if len(devices) > 1 {
		device, err = selectDeviceWithPrompt(devices)
		if err != nil {
			return err
		}
	} else {
		device = devices[0]
	}

	path, err := adb.Screencap(device)
	if err != nil {
		return err
	}

	if err := adb.Pull(device, path); err != nil {
		return err
	}

	fmt.Println("screencap is success.")

	return nil
}
