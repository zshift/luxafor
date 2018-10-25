package luxafor

import (
	"time"

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

func (lux Luxafor) sendCommand(command byte, led LED, r, g, b, speed uint8) (err error) {
	info := lux.deviceInfo
	device, err := info.Open()
	if err != nil {
		return errors.Wrap(err, "open lux")
	}

	defer func() { _ = device.Close() }() // Best effort.

	// Sets specified LED to RGB.
	if _, err := device.Write([]byte{command, byte(led), r, g, b}); err != nil {
		return errors.Wrap(err, "device write")
	}
	return nil
}

// Solid turns the specified luxafor into a solid RGB color.
func (lux Luxafor) Solid(r, g, b uint8) (err error) {
	return lux.Set(All, r, g, b)
}

// Set sets a golux.LED to the specific RGB value.
func (lux Luxafor) Set(led LED, r, g, b uint8) (err error) {
	return lux.sendCommand(static, led, r, g, b, 0) // speed isn't used
}

// Sets sets multiple golux.LED to the specific RGB value.
func (lux Luxafor) Sets(leds []LED, r, g, b uint8) (err error) {
	for _, led := range leds {
		if err := lux.Set(led, r, g, b); err != nil {
			return errors.Wrap(err, "set led")
		}
	}
	return nil
}

// Fade sets the led to rgb at speed.
func (lux Luxafor) Fade(led LED, r, g, b, speed uint8) (err error) {
	return lux.sendCommand(fade, led, r, g, b, speed)
}

// Police look like da popo
func (lux Luxafor) Police(loops int) (err error) {
	for i := 0; i < loops; i++ {
		lux.sendCommand(fade, FrontAll, 255, 0, 0, 255)
		lux.sendCommand(fade, BackAll, 0, 0, 255, 255)
		time.Sleep(500 * time.Millisecond)
		lux.sendCommand(fade, FrontAll, 0, 0, 255, 255)
		lux.sendCommand(fade, BackAll, 255, 0, 0, 255)
		time.Sleep(500 * time.Millisecond)
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
