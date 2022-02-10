package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/yasukotelin/adbeem/adb"
	"github.com/yasukotelin/adbeem/ffmpeg"
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

	output, err := screenRecord(cli, adb)
	if err != nil {
		return err
	}

	if cli.Bool("gif") {
		convertToGif(cli, output)
	}

	return nil
}

func screenRecord(cli *cli.Context, adb *adb.Adb) (string, error) {
	remote, err := adb.Screenrecord()
	if err != nil {
		return "", err
	}

	output := cli.String("output")
	if output == "" {
		output = filepath.Base(remote)
	}
	if err := adb.Pull(remote, output); err != nil {
		return "", err
	}

	if err := adb.Rm(remote); err != nil {
		return "", err
	}

	fmt.Println("screenrecord is success")

	return output, nil
}

func convertToGif(cli *cli.Context, input string) error {
	fmt.Println("Convert to gif...")

	ffmpeg := &ffmpeg.Ffmpeg{
		Input: input,
	}

	quality := cli.String("gif-quality")
	fps := cli.String("gif-fps")
	size := cli.String("gif-size")
	orientation := cli.String("gif-orientation")

	if err := ffmpeg.ConvToGif(quality, fps, size, orientation); err != nil {
		return err
	}

	fmt.Println("Successful gif conversion")

	return nil
}
