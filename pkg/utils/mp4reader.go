package utils

import (
	"bytes"
	"os/exec"
	"time"
)

// ExtractImage, stolen from https://stackoverflow.com/questions/57838025/get-frame-from-video-bytes
// Great idea!
func ExtractImage(fileBytes []byte){

	// command line args, path, and command
	command := "ffmpeg"
	frameExtractionTime := "0:00:05.000"
	vframes := "1"
	qv := "2"
	output := `c:\temp\mp4processing` + time.Now().Format(time.Kitchen) + ".jpg"

	cmd := exec.Command(command,
		"-ss", frameExtractionTime,
		"-i", "-",  // to read from stdin
		"-vframes", vframes,
		"-q:v", qv,
		output)

	cmd.Stdin = bytes.NewBuffer(fileBytes)

	// run the command and don't wait for it to finish. waiting exec is run
	// ignore errors for examples-sake
	_ = cmd.Start()
	_ = cmd.Wait()
}

