// Command relay is a evdev example demonstrating relaying events from an input
// device to a uinput device.
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kenshaw/evdev"
)

var (
	flagNode = flag.String("node", "", "input device")
)

func main() {
	flag.Parse()

	// open input
	d, err := evdev.OpenFile(*flagNode)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	log.Printf("opened %s [%s] from %s", d.Name(), d.Serial(), *flagNode)

	// create uinput device
	u, err := evdev.NewUserInput(
		0644,
		evdev.WithName(d.Name()+" relay"),
		evdev.WithID(d.ID()),
		evdev.WithTypesFromEvdev(d),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer u.Close()

	log.Printf("created %q [%s] for %s: %s", d.Name(), d.Serial(), *flagNode, u.Path())

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start polling and relaying
	in := d.Poll(ctxt)
	go func() {
		for {
			select {
			case <-ctxt.Done():
				return

			case event := <-in:
				if event == nil {
					return
				}
				log.Printf("<- %+v", event)
				go u.Send(*event)
			}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("received signal: %s", <-sig)
}
