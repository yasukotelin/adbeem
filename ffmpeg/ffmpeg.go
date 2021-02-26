package ffmpeg

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Ffmpeg struct {
	Input  string
	Output string
}

func (ffmpeg *Ffmpeg) ExistsCommand() bool {
	cmd := exec.Command("ffmpeg", "-h")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func (ffmpeg *Ffmpeg) ConvToGif(fps string) error {
	output := ffmpeg.Output
	if output == "" {
		output = ffmpeg.createOutputFromInput()
	}
	cmd := exec.Command("ffmpeg", "-i", ffmpeg.Input, "-vf", "[0:v] fps="+fps+",scale=320:-1,split [a][b];[a] palettegen [p];[b][p] paletteuse", output)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (ffmpeg *Ffmpeg) createOutputFromInput() string {
	ext := filepath.Ext(ffmpeg.Input)
	return strings.Replace(ffmpeg.Input, ext, ".gif", 1)
}
