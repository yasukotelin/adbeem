package cmd

import "github.com/manifoldco/promptui"

func selectDeviceWithPrompt(devices []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select a device",
		Items: devices,
		Size:  5,
	}
	_, device, err := prompt.Run()

	return device, err
}
