package luxafor

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
func (lux Luxafor) Solid(r, g, b uint8) (err error) {
	return lux.SetLED(All, r, g, b)
}

// SetLED sets a luxafor.LED to the specific RGB value.
func (lux Luxafor) SetLED(led LED, r, g, b uint8) (err error) {
	info := lux.deviceInfo
	device, err := info.Open()
	if err != nil {
		return errors.Wrap(err, "open lux")
	}

	defer func() { _ = device.Close() }() // Best effort.

	// Sets specified LED to RGB.
	if _, err := device.Write([]byte{static, byte(led), r, g, b}); err != nil {
		return errors.Wrap(err, "device write")
	}
	return nil
}

// SetLEDs sets multiple luxafor.LED to the specific RGB value.
func (lux Luxafor) SetLEDs(leds []LED, r, g, b uint8) (err error) {
	for _, led := range leds {
		if err := lux.SetLED(led, r, g, b); err != nil {
			return errors.Wrap(err, "set led")
		}
	}
	return nil
}

// Off turns off the luxafor.
func (lux Luxafor) Off() (err error) {
	info := lux.deviceInfo
	device, err := info.Open()
	if err != nil {
		return errors.Wrap(err, "open lux")
	}

	defer func() { _ = device.Close() }() // Best effort.

	// Turns off the leds.
	if _, err := device.Write([]byte{static, byte(All), 0, 0, 0}); err != nil {
		return errors.Wrap(err, "device write")
	}
	return nil
}
