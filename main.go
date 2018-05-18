package main

import (
	"fmt"

	"github.com/karalabe/hid"
)

const (
	vendorID uint16 = 0x04d8
	deviceID uint16 = 0xf372

	static byte = 1
	fade   byte = 2
	strobe byte = 3
	wave   byte = 4
	pattrn byte = 6

	luxafor     byte = 1
	random1     byte = 2
	random2     byte = 3
	random3     byte = 4
	random4     byte = 6
	random5     byte = 7
	police      byte = 5
	rainbowWave byte = 8

	frontTop    byte = 1
	frontMiddle byte = 2
	frontBottom byte = 3
	backTop     byte = 4
	backMiddle  byte = 5
	backBottom  byte = 6
	frontAll    byte = 65
	backAll     byte = 66
	all         byte = 255
)

func main() {
	infos := hid.Enumerate(vendorID, deviceID)
	fmt.Printf("Found %d Luxaforas attached.\n", len(infos))
	for i, info := range infos {
		fmt.Printf("Connecting to device  %d\n", i)

		device, err := info.Open()
		if err != nil {
			fmt.Printf("Failed to open device %d\n", i)
			continue
		}

		// sets device to solid white
		device.Write([]byte{static, all, 255, 255, 255})
	}
}
