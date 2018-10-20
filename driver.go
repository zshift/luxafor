package golux

import (
	"github.com/karalabe/hid"
	"github.com/pkg/errors"
)

// Luxafor is used to access the devices.
type Luxafor struct {
	deviceInfo hid.DeviceInfo
}

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

// Enumerate returns a slice of attached Luxafors
func Enumerate() []Luxafor {
	infos := hid.Enumerate(vendorID, deviceID)
	luxs := make([]Luxafor, len(infos))
	for _, info := range infos {
		lux := Luxafor{
			deviceInfo: info,
		}
		luxs = append(luxs, lux)
	}

	return luxs
}

// Solid turns the specified luxafor into a solid RGB color.
func (lux *Luxafor) Solid(r, g, b uint8) (err error) {
	info := lux.deviceInfo
	device, err := info.Open()
	defer func() { _ = device.Close() }() // Best effort.
	if err != nil {
		return errors.Wrap(err, "open lux")
	}

	// sets device to solid color
	device.Write([]byte{static, all, r, g, b})
	return nil
}
