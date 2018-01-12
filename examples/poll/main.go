// Command poll is a evdev example demonstrating how to poll a input device for
// events.
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

	// open
	d, err := evdev.OpenFile(*flagNode)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

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
