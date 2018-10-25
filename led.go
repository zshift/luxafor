package luxafor

// LED represents the location of an individual LED to control on the Luxafor.
type LED byte

// Exported locations that represent individual--and sets of--LEDs on the Luxafor.
const (
	FrontTop    LED = 1
	FrontMiddle LED = 2
	FrontBottom LED = 3
	BackTop     LED = 4
	BackMiddle  LED = 5
	BackBottom  LED = 6
	FrontAll    LED = 65
	BackAll     LED = 66
	All         LED = 255
)
