package luxafor

import (
	"testing"
	"time"
)

func finishTest(lux Luxafor) {
	time.Sleep(250 * time.Millisecond)
	lux.Off()
}

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
	defer finishTest(lux)

	if err := lux.Solid(r, g, b); err != nil {
		t.Errorf("Failed to turn solid %s.", color)
	}
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
	defer finishTest(lux)

	lux.Solid(255, 255, 255)
}

func TestSet(t *testing.T) {
	luxs := Enumerate()
	if len(luxs) == 0 {
		t.Log("No attached devices. Aborting test.")
		return
	}

	lux := luxs[1]
	defer finishTest(lux)

	lux.Set(FrontAll, 255, 0, 0)
	time.Sleep(250 * time.Millisecond)
	lux.Set(BackAll, 0, 255, 0)
}

func TestSetLEDs(t *testing.T) {
	luxs := Enumerate()
	if len(luxs) == 0 {
		t.Log("No attached devices. Aborting test.")
		return
	}

	lux := luxs[1]
	defer finishTest(lux)

	if err := lux.Sets([]LED{FrontAll, BackMiddle, BackTop}, 255, 0, 0); err != nil {
		t.Error(err.Error())
	}

	time.Sleep(250 * time.Millisecond)

	if err := lux.Sets([]LED{FrontMiddle, FrontBottom}, 123, 53, 98); err != nil {
		t.Error(err.Error())
	}
}

func TestFade(t *testing.T) {
	luxs := Enumerate()
	if len(luxs) == 0 {
		t.Log("No attached devices. Aborting test.")
		return
	}

	lux := luxs[1]
	defer finishTest(lux)

	if err := lux.Fade(FrontAll, 0, 0, 255, 255); err != nil {
		t.Error(err.Error())
		t.Log("asdf")
	}
}

func TestPolice(t *testing.T) {
	luxs := Enumerate()
	if len(luxs) == 0 {
		t.Log("No attached devices. Aborting test.")
		return
	}

	lux := luxs[1]
	defer finishTest(lux)
	lux.Police(100)
}
