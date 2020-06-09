package cmd

import (
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/yasukotelin/adbeem/adb"
)

func selectDeviceWithPrompt(adb *adb.Adb) (string, error) {
	devices, err := adb.GetDevices()
	if err != nil {
		return "", err
	}

	var device string
	switch len(devices) {
	case 0:
		return "", errors.New("No connected devices")
	case 1:
		device = devices[0]
	default:
		prompt := promptui.Select{
			Label: "Select a device",
			Items: devices,
			Size:  5,
		}
		_, device, err = prompt.Run()
		if err != nil {
			return "", err
		}
	}

	return device, err
}
