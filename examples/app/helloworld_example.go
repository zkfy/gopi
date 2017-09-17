/*
	Go Language Raspberry Pi Interface
	(c) Copyright David Thorpe 2016-2017
	All Rights Reserved
	Documentation http://djthorpe.github.io/gopi/
	For Licensing and Usage information, please see LICENSE.md
*/

// The canonical hello world example demonstrates printing
// hello world and then exiting. Here we use the 'generic'
// set of modules which provide generic system services
package main

import (
	"fmt"
	"os"

	"github.com/djthorpe/gopi"
	_ "github.com/djthorpe/gopi/sys/logger"
)

////////////////////////////////////////////////////////////////////////////////

func helloWorld(app *gopi.AppInstance, done chan struct{}) error {
	// If -name argument is used then use that, else output generic message
	if name, exists := app.AppFlags.GetString("name"); exists {
		fmt.Println("Hello,", name)
	} else {
		fmt.Println("Hello, World")
	}
	done <- gopi.DONE
	return nil
}

func registerFlags(config *gopi.AppConfig) {
	// Register the -name flag
	config.AppFlags.FlagString("name", "", "Your name")
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	config := gopi.NewAppConfig()
	registerFlags(&config)
	if app, err := gopi.NewAppInstance(config); err != nil {
		// Check to see if -help has been triggered
		if err != gopi.ErrHelp {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
	} else if err := app.Run(helloWorld); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
