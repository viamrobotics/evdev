// Command iter is a evdev example demonstrating how to iterate input devices
// and read their configuration.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/kenshaw/evdev"
)

var (
	flagVerbose = flag.Bool("v", false, "toggle verbose")
	flagRumble  = flag.Bool("rumble", true, "toggle rumble")
)

func main() {
	flag.Parse()

	devs, err := filepath.Glob("/dev/input/event*")
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(devs, func(i, j int) bool {
		a, _ := strconv.Atoi(strings.TrimPrefix(devs[i], "/dev/input/event"))
		b, _ := strconv.Atoi(strings.TrimPrefix(devs[j], "/dev/input/event"))
		return a <= b
	})
	for i, n := range devs {
		d, err := evdev.OpenFile(n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: unable to open %s: %v\n", n, err)
			continue
		}
		defer d.Close()

		id := d.ID()
		var e string
		if s := d.Serial(); s != "" {
			e = " [" + s + "]"
		}

		if i != 0 {
			fmt.Fprintln(os.Stdout)
		}

		effects, effectStatuses := d.EffectTypes(), d.EffectStatusTypes()

		// print information about the device
		fmt.Fprintf(os.Stdout, "%s\n", n)
		fmt.Fprintf(os.Stdout, "  Device: %s\n", d.Name()+e)
		fmt.Fprintf(os.Stdout, "  Path: %s\n", d.Path())
		fmt.Fprintf(os.Stdout, "  Bus: %s, Vendor: 0x%04x, Product: 0x%04x, Version: 0x%0x\n", id.BusType, id.Vendor, id.Product, id.Version)
		if *flagVerbose {
			fmt.Fprintf(os.Stdout, "  Keys/Buttons: %s\n",
				format(d.KeyTypes()),
			)
			fmt.Fprintf(os.Stdout, "  Absolute Axis: %s, Relative Axis: %s\n",
				format(d.AbsoluteTypes()),
				format(d.RelativeTypes()),
			)
			fmt.Fprintf(os.Stdout, "  Effects: %s, Effect Statuses: %v\n",
				format(effects),
				format(effectStatuses),
			)
			fmt.Fprintf(os.Stdout, "  Switches: %s, LEDs: %s, Sounds: %s, Power: %s\n",
				format(d.SwitchTypes()),
				format(d.LEDTypes()),
				format(d.SoundTypes()),
				format(d.PowerTypes()),
			)
		}

		/*if len(effects) != 0 {
			if err := d.Lock(); *flagRumble && err == nil {
				// normalize gain
				d.EffectGainSet(75)

				effect := d.EffectNew(
					evdev.RumbleEffect(
						0,
						0,
					),
				)
				d.Unlock()
			}
		}*/
	}
}

// format does reflect magic since all types are of map[<type>]bool where each
// <type> satisfies fmt.Stringer interface.
func format(x interface{}) string {
	mk := reflect.ValueOf(x).MapKeys()
	if len(mk) == 0 {
		return "<none>"
	}

	var i int
	ityp := reflect.TypeOf(i)
	keys := make([]int, len(mk))
	m := make(map[int]string, len(mk))
	for n, k := range mk {
		z := int(k.Convert(ityp).Int())
		keys[n] = z
		o := fmt.Sprintf("%v", k)
		if strings.Contains(o, "(") {
			o = fmt.Sprintf("0x%03x", z)
		}
		m[z] = o
	}
	sort.Slice(keys, func(a, b int) bool { return keys[a] < keys[b] })
	s := "["
	for i, k := range keys {
		if i != 0 {
			s += ", "
		}
		s += m[k]
	}
	return s + "]"
}
