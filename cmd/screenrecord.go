package cmd

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/yasukotelin/adbeem/adb"
)

func Screenrecord(cli *cli.Context) error {
	adb := adb.New()

	if !adb.ExistsCommand() {
		return errors.New("adb command not found")
	}

	device, err := selectDeviceWithPrompt(adb)
	if err != nil {
		return err
	}

	path, err := adb.Screenrecord(device)
	if err != nil {
		return err
	}

	output := cli.String("output")
	if err := adb.Pull(device, path, output); err != nil {
		return err
	}

	if err := adb.Rm(device, path); err != nil {
		return err
	}

	fmt.Println("screenrecord is success")

	return nil
}
