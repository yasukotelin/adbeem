# adbeem

adbeem is a CLI tool to capture or record the screen of Android with adb.

## Features

- Easy capture screen
- Easy record screen
- Multi devices. You can select a device with prompt ui.

## Install

```
go get github.com/yasukotelin/adbeem
```

## Usage

Show help.

```
$ adbeem
NAME:
   adbeem - adbeem is a CLI tool to capture or record the screen of Android with adb

USAGE:
   adbeem [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   screencap     capture screen
   screenrecord  records screen
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

Take a screen capture.

```
adbeem screencap
```

Captures the screen of the connected device and outputs a png file to the execution path.

```
adbeem screenrecord
```

Records the screen of the connected device and outputs a mp4 file to the execution path.

## Author

yasukotelin
