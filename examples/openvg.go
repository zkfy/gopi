/*
	Go Language Raspberry Pi Interface
	(c) Copyright David Thorpe 2016
	All Rights Reserved

	For Licensing and Usage information, please see LICENSE.md
*/

// This example demonstrates the use of OpenVG for vector graphics
package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
)

import (
	rpi "../device/rpi"
	openvg "../openvg"
	util "../util"
)

////////////////////////////////////////////////////////////////////////////////

var (
	flagDisplay = flag.String("display", "lcd", "Display")
)
////////////////////////////////////////////////////////////////////////////////

func main() {
	// Flags
	flag.Parse()

	// Create logger
	logger := new(util.StderrLogger)

	// Retrieve display
	display, ok := rpi.Displays[*flagDisplay]
	if !ok {
		fmt.Fprintln(os.Stderr, "Error: Invalid display name")
		return
	}

	// Open up the RaspberryPi interface
	pi, err := rpi.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		return
	}
	defer pi.Close()

	// VideoCore
	vc, err := pi.NewVideoCore(display)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		return
	}
	defer vc.Close()

	// OpenVG
	gfx,err := openvg.Open(&rpi.OpenVG{ VideoCore: vc, Logger: logger })
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		return
	}
	defer gfx.Close()

	fmt.Println("Press any key to exit...")
	reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')

}