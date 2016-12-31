/*
    GOPI Go Language Raspberry Pi Interface
	(c) Copyright David Thorpe 2016
	All Rights Reserved

	For Licensing information, please see LICENSE.md
	For Documentation, see http://djthorpe.github.io/gopi/
*/

// This example outputs a table of detected input devices, their types
// and other information about them.
package main

import (
	"fmt"
	"os"
)

import (
	app "github.com/djthorpe/gopi/app"
	hw "github.com/djthorpe/gopi/hw"
)

////////////////////////////////////////////////////////////////////////////////

func MyRunLoop(app *app.App) error {

	// Opens all devices
	app.Logger.Info("input=%v", app.Input)
	devices, err := app.Input.OpenDevicesByName("", hw.INPUT_TYPE_ANY, hw.INPUT_BUS_ANY)
	if err != nil {
		return err
	}

	format := "%-30s %-25s %-25s\n"
	fmt.Printf(format, "Name", "Type", "Bus")
	fmt.Printf(format, "------------------------------", "-------------------------", "-------------------------")

	for _, device := range devices {
		fmt.Printf(format, device.GetName(), device.GetType(), device.GetBus())
	}

	// Return success
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	// Create the config
	config := app.Config(app.APP_INPUT)

	// Flags
	config.FlagSet.FlagString("type", "", "Input type (any, mouse, keyboard, joystick, touchscreen)")
	config.FlagSet.FlagString("bus", "", "Input bus (any, usb, bluetooth)")
	config.FlagSet.FlagString("name", "", "Name of input device")
	config.FlagSet.FlagBool("watch", false, "Watch for events from devices until CTRL+C is pressed")

	// Create the application
	myapp, err := app.NewApp(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	defer myapp.Close()

	// Run the application
	if err := myapp.Run(MyRunLoop); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
}
