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

	adb.Device = device

	if err = screenRecord(cli, adb); err != nil {
		return err
	}

	if cli.Bool("gif") {
		convertToGif(cli)
	}

	return nil
}

func screenRecord(cli *cli.Context, adb *adb.Adb) error {
	path, err := adb.Screenrecord()
	if err != nil {
		return err
	}

	output := cli.String("output")
	if err := adb.Pull(path, output); err != nil {
		return err
	}

	if err := adb.Rm(path); err != nil {
		return err
	}

	fmt.Println("screenrecord is success")

	return nil
}

func convertToGif(cli *cli.Context) error {
	fmt.Println("Convert to gif...")

	// TODO ffmpeg command check
	return nil
}
