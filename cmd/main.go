package main

import (
	"fmt"

	"github.com/zshift/luxafor"
)

func main() {
	luxs := luxafor.Enumerate()
	if len(luxs) == 0 {
		fmt.Println("No attached devices. Exiting")
		return
	}

	fmt.Printf("Found %v devices.", len(luxs))

	lux := luxs[1]
	err := lux.Solid(255, 0, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
}
