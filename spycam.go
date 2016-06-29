package main

import "github.com/blackjack/webcam"
import (
	"fmt"
	"io/ioutil"
	"os"
)

type FrameSizes []webcam.FrameSize

func findMjpegFormat(cam *webcam.Webcam) (*webcam.PixelFormat, error) {
	format_desc := cam.GetSupportedFormats()

	for format, formatName := range format_desc {
		if(formatName == "MJPEG") {
			return &format, nil
		}
	}

	return nil, fmt.Errorf("Could not find MJPEG format")
}

func selectHighestResolution(cam *webcam.Webcam, format webcam.PixelFormat) webcam.FrameSize {
	frameSizes := FrameSizes(cam.GetSupportedFrameSizes(format))
	var highestResolution *webcam.FrameSize
	for _, frameSize := range frameSizes {
		if highestResolution == nil || frameSize.MaxWidth > highestResolution.MaxWidth {
			highestResolution = &frameSize
		}
	}

	return *highestResolution
}

func saveFrame(frame []byte, i uint32) {
	fileName := fmt.Sprintf("/home/marin/saved_frames/frame_%05d.jpeg", i)
	err := ioutil.WriteFile(fileName, frame, 0644)
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}
}

func main() {
	cam, err := webcam.Open("/dev/video0")
	if err != nil {
		panic(err.Error())
	}
	defer cam.Close()

	mjpegFormat, err := findMjpegFormat(cam)
	if(err != nil) {
		panic(err.Error())
	}

	frameSize := selectHighestResolution(cam, *mjpegFormat)
	_, _, _, err1 := cam.SetImageFormat(*mjpegFormat, frameSize.MaxWidth, frameSize.MaxHeight)
	if err1 != nil {
		panic(err1.Error())
	}

	fmt.Println("Capturing a frame!")

	err = cam.StartStreaming()
	if err != nil {
		panic(err.Error())
	}

	timeout := uint32(5) //5 seconds
	for frameNum := uint32(0); true; frameNum++ {
		err = cam.WaitForFrame(timeout)

		switch err.(type) {
		case nil:
		case *webcam.Timeout:
			fmt.Fprint(os.Stderr, err.Error())
			continue
		default:
			panic(err.Error())
		}

		frame, err := cam.ReadFrame()
		if len(frame) != 0 {
			print(".")
			saveFrame(frame, frameNum)
		} else if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("Done!")
}
