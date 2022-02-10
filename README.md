# adbeem

![version](https://img.shields.io/badge/version-1.6.1-brightgreen)
[![GoDoc](https://godoc.org/github.com/yasukotelin/adbeem?status.svg)](https://godoc.org/github.com/yasukotelin/adbeem)
[![Go Report Card](https://goreportcard.com/badge/github.com/yasukotelin/adbeem)](https://goreportcard.com/report/github.com/yasukotelin/adbeem)

adbeem is a CLI tool using Android "adb" command more easily.

## Index

* [Requirement](#requirement)
* [Features](#features)
* [Install](#install)
* [Update](#update)
* [Usage](#usage)
  * [screencap](#screencap)
  * [screenrecord](#screenrecord)
    * [Convert to gif](#convert-to-gif)
  * [deeplink](#deeplink)
* [Author](#author)

## Requirement

- [adb](https://developer.android.com/studio/command-line/adb)
- [go](https://golang.org/doc/install) >= 13

If you use to convert to gif from mp4 screen capture.
- ffmpeg

## Features

- Easy capture screen（JPEG, PNG）.
- Easy record screen（MP4）.
- Easy send deep link to device.

Multi devices. You can select a device with prompt ui.

## Install

You need to install [adb](https://developer.android.com/studio/command-line/adb) before install adbeem.

If you use to convert to gif from mp4 screen capture, you need to install ffmpeg.

And you can install adbeem with go get.

```
go install github.com/yasukotelin/adbeem@latest
```

## Usage

Show help.

```
╰─>$ adbeem --help
NAME:
   adbeem - adbeem is a CLI tool using Android adb command more easily.

USAGE:
   adbeem [global options] command [command options] [arguments...]

VERSION:
   1.6.0

COMMANDS:
   screencap, screenshot  capture the screen
   screenrecord           records the screen
   deeplink               send deeplink
   help, h                Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### screencap

```
adbeem screencap
# or
adbeem screenshot
```

If you use --output (or -o), you can specify output path.<br>
Default output path is current directory and like this format `adbeem-20200611010217.png`.

```
adbeem screencap -o ~/Desktop/output.png
```

### screenrecord

Records the screen on device.(Max 3 minute)

```
adbeem screenrecord
```

screenrecord command can also use --output flag.

```
adbeem screenrecord -o ~/Desctop/demo.mp4
```

#### Convert to gif

If you want gif, you can use --gif(-g) option.

```
adbeem screenrecord -g
```

adbeem outputs mp4 and gif file.

--gif-quality option. middle or high.
default is middle.
If you use high option, output file is bigger than middle one.

```
adbeem screenrecord -g --gif-quality high
```

--gif-size, gif size option. (default is 320)
size is width pixel. If you use `--gif-orientation landscape`, size is height pixel.

```
adbeem screenrecord -g --gif-size 800
```

--gif-orientation, portlate or landscape. (default is portlait)

```
adbeem screenrecord -g --gif-orientation landscape
```

--gif-fps, gif fps option. (default is 20)

```
abdeem screenrecord -g --gif-fps 30
```

### deeplink

Send deep link to connected device.

```
adbeem deeplink -u "https://github.com"
```

## Author

yasukotelin
