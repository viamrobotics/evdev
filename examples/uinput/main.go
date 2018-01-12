// Command uinput is a evdev example demonstrating creating a user input
// device.
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kenshaw/evdev"
)

func main() {
	u, err := evdev.NewUserInput(
		0644,
		evdev.WithID(evdev.ID{
			BusType: evdev.BusUSB,
			Vendor:  0x01,
			Product: 0x02,
			Version: 0x0a0b,
		}),
		evdev.WithName("test input device"),
		evdev.WithPath("something"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("created: %s", u.Path())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("received interrupt: %s", <-sig)
}
