// Command led is a evdev example demonstrating how to send events (LED) to an
// input device.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/kenshaw/evdev"
)

const Timeout = 200 * time.Millisecond

func main() {
	node := parseArgs()

	// Create and open our device.
	dev, err := evdev.OpenFile(node)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	// Make sure it is closed once we are done.
	defer dev.Close()

	// Turn off the Capslock, NumLock and ScrollLock LEDs.
	var ev evdev.Event
	ev.Type = evdev.EvLed
	ev.Value = 0

	ev.Code = evdev.LedCapsLock
	dev.Outbox <- ev

	ev.Code = evdev.LedNumLock
	dev.Outbox <- ev

	ev.Code = evdev.LedScrollLock
	dev.Outbox <- ev

	// Once every 200 milliseconds, toggle one of the LEDs.
	// Or exit if we receive an exit signal.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	ev.Code = evdev.LedNumLock

	for {
		select {
		case <-signals:
			return

		case <-time.After(Timeout):
			// Turn off previous LED
			ev.Value = 0
			dev.Outbox <- ev

			// Turn on the next one.
			ev.Code = (ev.Code + 1) & 3
			ev.Value = 1
			dev.Outbox <- ev
		}
	}
}

func parseArgs() string {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s <node>\n", os.Args[0])
		os.Exit(1)
	}

	return flag.Args()[0]
}
