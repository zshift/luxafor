package golux

import (
	"testing"
	"time"
)

func TestEnumerate(t *testing.T) {
	luxs := Enumerate()
	t.Logf("Number of attached devices: %v", len(luxs))
}

func testSolid(t *testing.T, r, g, b uint8, color string) {
	luxs := Enumerate()
	if len(luxs) == 0 {
		t.Log("No attached devices. Aborting test.")
	}

	// TODO: 0 always fails. Need to investigate why.
	lux := luxs[1]

	defer lux.Off()

	err := lux.Solid(r, g, b)
	if err != nil {
		t.Errorf("Failed to turn solid %s.", color)
	} else {
		defer lux.Off()
	}
	time.Sleep(250 * time.Millisecond)
}

func TestWhite(t *testing.T) {
	testSolid(t, 255, 255, 255, "white")
}

func TestRed(t *testing.T) {
	testSolid(t, 255, 0, 0, "red")
}

func TestGreen(t *testing.T) {
	testSolid(t, 0, 255, 0, "green")
}

func TestBlue(t *testing.T) {
	testSolid(t, 0, 0, 255, "blue")
}

func TestCyan(t *testing.T) {
	testSolid(t, 0, 255, 255, "cyan")
}

func TestMagenta(t *testing.T) {
	testSolid(t, 255, 0, 255, "magenta")
}

func TestYellow(t *testing.T) {
	testSolid(t, 255, 255, 0, "yellow")
}

func TestOff(t *testing.T) {
	luxs := Enumerate()
	if len(luxs) == 0 {
		t.Log("No attached devices. Aborting test.")
		return
	}

	lux := luxs[1]

	lux.Solid(255, 255, 255)
	time.Sleep(500 * time.Millisecond)
	lux.Off()
}

func TestSetLED(t *testing.T) {
	luxs := Enumerate()
	if len(luxs) == 0 {
		t.Log("No attached devices. Aborting test.")
		return
	}

	lux := luxs[1]
	defer lux.Off()

	lux.SetLED(FrontAll, 255, 0, 0)
	lux.SetLED(BackAll, 0, 255, 0)
}

func TestSetLEDs(t *testing.T) {
	luxs := Enumerate()
	if len(luxs) == 0 {
		t.Log("No attached devices. Aborting test.")
		return
	}

	lux := luxs[1]
	defer lux.Off()

	err := lux.SetLEDs([]LED{FrontAll, BackMiddle, BackTop}, 255, 0, 0)
	if err != nil {
		t.Error(err.Error())
	}
}
