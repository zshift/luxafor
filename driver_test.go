package golux

import (
	"testing"
	"time"
)

func TestEnumerate(t *testing.T) {
	luxs := Enumerate()
	t.Logf("Number of attached devices: %v", len(luxs))
}

func TestSolid(t *testing.T) {
	luxs := Enumerate()
	if len(luxs) == 0 {
		t.Log("No attached devices. Aborting test.")
	}

	lux := luxs[1]
	// black (aka off)
	defer lux.Solid(0, 0, 0)

	// white
	lux.Solid(255, 255, 255)
	time.Sleep(1 * time.Second)

	// red
	lux.Solid(255, 0, 0)
	time.Sleep(1 * time.Second)

	// green
	lux.Solid(0, 255, 0)
	time.Sleep(1 * time.Second)

	// blue
	lux.Solid(0, 0, 255)
	time.Sleep(1 * time.Second)

	// cyan
	lux.Solid(0, 255, 255)
	time.Sleep(1 * time.Second)

	// magenta
	lux.Solid(255, 0, 255)
	time.Sleep(1 * time.Second)

	// yellow
	lux.Solid(255, 255, 0)
	time.Sleep(1 * time.Second)
}
