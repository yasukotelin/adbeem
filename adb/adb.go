package adb

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type Adb struct {
	Device string
}

func New() *Adb {
	return &Adb{}
}

func (adb *Adb) ExistsCommand() bool {
	cmd := exec.Command("adb", "--version")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func (adb *Adb) GetDevices() ([]string, error) {
	cmd := exec.Command("adb", "devices")
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	rows := strings.Split(string(out), "\n")[1:]

	devices := make([]string, 0, len(rows))
	for _, row := range rows {
		if row == "" {
			continue
		}
		devices = append(devices, strings.SplitN(row, "\t", 2)[0])
	}

	return devices, nil
}

func (adb *Adb) Screencap() (string, error) {
	path := "/sdcard/" + getFileName() + ".png"

	cmd := adb.newCmd("shell", "screencap", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return path, nil
}

func (adb *Adb) Screenrecord() (string, error) {
	path := "/sdcard/" + getFileName() + ".mp4"

	cmd := adb.newCmd("shell", "screenrecord", path)
	if err := cmd.Start(); err != nil {
		return "", err
	}

	fmt.Println("Start recording...")
	fmt.Print("(press any button to stop) ")

	go func() {
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()

		// Send to stop process signal
		cmd.Process.Signal(syscall.SIGINT)
	}()

	cmd.Wait()

	// Wait to stop screenrecord process.
	for {
		rows, err := adb.PsTarget("screenrecord")
		if err != nil {
			return "", err
		}

		if len(rows) == 0 {
			break
		}
	}

	fmt.Println("stop record")

	return path, nil
}

func getFileName() string {
	t := time.Now()
	const layout = "20060102150405"
	return "adbeem-" + t.Format(layout)
}

func (adb *Adb) Pull(remote string, local string) error {
	cmd := adb.newCmd("pull", remote, local)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (adb *Adb) Rm(path string) error {
	cmd := adb.newCmd("shell", "rm", path)
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (adb *Adb) Ps() ([]string, error) {
	cmd := adb.newCmd("shell", "ps")
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	rows := strings.Split(string(out), "\n")

	// Triming last row
	rows = rows[:len(rows)-1]

	return rows, nil
}

func (adb *Adb) PsTarget(target string) ([]string, error) {
	psList, err := adb.Ps()
	if err != nil {
		return nil, err
	}

	var result []string
	for _, row := range psList {
		if strings.Contains(row, target) {
			result = append(result, row)
		}
	}

	return result, nil
}

func (adb *Adb) SendDeepLink(url string) error {
	cmd := adb.newCmd("shell", "am", "start", "-a", "android.intent.action.VIEW", "-d", url)
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (adb *Adb) newCmd(args ...string) *exec.Cmd {
	arg := make([]string, 0, 10)
	if adb.Device != "" {
		arg = append(arg, "-s", adb.Device)
	}
	arg = append(arg, args...)

	return exec.Command("adb", arg...)
}
