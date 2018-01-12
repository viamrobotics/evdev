// Command uinput is a evdev example demonstrating how to create a virtual user
// input ("uinput") device and sending/receiving events for it.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/kenshaw/evdev"
)

var (
	flagNode = flag.String("node", "", "input device")
)

func main() {
	flag.Parse()

	// create
	d, err := evdev.NewUserDevice()
	if err != nil {
		log.Fatal(err)
	}

	// start polling
	ch := d.Poll(context.Background())

loop:
	for {
		select {
		case event := <-ch:
			// channel closed
			if event == nil {
				break loop
			}

			switch typ := event.Type.(type) {
			case evdev.KeyType:
				if typ == evdev.KeyQ {
					log.Printf("quitting")
					break loop
				}
				log.Printf("received key event: %+v", event)

			case evdev.AbsoluteType:
				log.Printf("received absolute axis event: %+v", event)
				log.Printf("   axis information: %+v", d.AbsoluteTypes()[typ])
			}
		}
	}
}
