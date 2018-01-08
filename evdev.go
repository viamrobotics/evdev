// Package evdev is a pure Go implementation of the Linux evdev API.
package evdev

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"sync"
	"unsafe"
)

const (
	// DefaultPollSize is the default number of events to poll.
	DefaultPollSize = 64
)

// Evdev represents an evdev device.
type Evdev struct {
	fd       *os.File
	pollSize int

	id        ID
	name      string
	path      string
	serial    string
	version   uint32
	effectMax int32

	events    map[EventType]bool
	syncs     map[SyncType]bool
	keys      map[KeyType]bool
	miscs     map[MiscType]bool
	absolutes map[AbsoluteType]Axis
	relatives map[RelativeType]bool
	switches  map[SwitchType]bool
	leds      map[LEDType]bool
	sounds    map[SoundType]bool
	repeats   map[RepeatType]bool
	effects   map[EffectType]bool
	powers    map[PowerType]bool
	effectss  map[EffectStatusType]bool

	out    chan Event
	cancel context.CancelFunc
}

// Open creates a device from the aready open file descriptor.
func Open(fd *os.File) *Evdev {
	return &Evdev{fd: fd}
}

// OpenFile opens device from the file path (ie, /dev/input/event*).
func OpenFile(path string) (*Evdev, error) {
	fd, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	return &Evdev{
		fd: fd,
	}, nil
}

// Close closes the underlying device file descriptor.
func (d *Evdev) Close() error {
	if d.cancel != nil {
		d.cancel()
	}
	if d.fd != nil {
		d.Unlock()
		err := d.fd.Close()
		d.fd = nil
		return err
	}
	return nil
}

// Lock attempts to gain exclusive access to the device.
//
// This means that we are the only ones receiving events from the device; other
// processes will not.
//
// This ability should be handled with care, especially when trying to lock
// keyboard access. If this is executed while we are running in something like
// X, this call will prevent X from receiving any and all keyboard events. All
// of them will only be sent to our own process. If we do not properly handle
// these key events, we may lock ourselves out of the system and a hard reset
// is required to restore it.
func (d *Evdev) Lock() error {
	return ioctl(d.fd.Fd(), _EVIOCGRAB, 1)
}

// Unlock releases a lock, previously obtained through Lock.
func (d *Evdev) Unlock() error {
	return ioctl(d.fd.Fd(), _EVIOCGRAB, 0)
}

// ID returns a device's identity information.
func (d *Evdev) ID() ID {
	var once sync.Once
	once.Do(func() {
		ioctl(d.fd.Fd(), _EVIOCGID, unsafe.Pointer(&d.id))
	})
	return d.id
}

// Name returns the name of the device.
func (d *Evdev) Name() string {
	var once sync.Once
	once.Do(func() {
		buf := make([]byte, 256)
		ioctl(d.fd.Fd(), _EVIOCGNAME(256), unsafe.Pointer(&buf[0]))
		if i := bytes.IndexByte(buf, 0); i != -1 {
			buf = buf[:i]
		}
		d.name = string(buf)
	})
	return d.name
}

// Path returns the physical path of the device. For example:
//
//    usb-00:01.2-2.1/input0
//
// To understand what this string is showing, you need to break it down into
// parts. `usb` means this is a physical topology from the USB system.
//
// `00:01.2` is the PCI bus information for the USB host controller (in this
// case, bus 0, slot 1, function 2).
//
// `2.1` shows the path from the root hub to the device. In this case, the
// upstream hub is plugged in to the second port on the root hub, and that
// device is plugged in to the first port on the upstream hub.
//
// `input0` means this is the first event interface on the device. Most
// devices have only one, but multimedia keyboards may present the normal
// keyboard on one interface and the multimedia function keys on a second
// interface.
func (d *Evdev) Path() string {
	var once sync.Once
	once.Do(func() {
		buf := make([]byte, 256)
		ioctl(d.fd.Fd(), _EVIOCGPHYS(256), unsafe.Pointer(&buf[0]))
		if i := bytes.IndexByte(buf, 0); i != -1 {
			buf = buf[:i]
		}
		d.path = string(buf)
	})
	return d.path
}

// Serial returns the unique serial for the device.
//
// Most devices do not have this and will return an empty string.
func (d *Evdev) Serial() string {
	var once sync.Once
	once.Do(func() {
		buf := make([]byte, 256)
		ioctl(d.fd.Fd(), _EVIOCGUNIQ(256), unsafe.Pointer(&buf[0]))
		if i := bytes.IndexByte(buf, 0); i != -1 {
			buf = buf[:i]
		}
		d.serial = string(buf)
	})
	return d.serial
}

// Version returns the major, minor, and revision of the device driver.
func (d *Evdev) Version() (int, int, int) {
	var once sync.Once
	once.Do(func() {
		ioctl(d.fd.Fd(), _EVIOCGVERSION, unsafe.Pointer(&d.version))
	})
	return int(d.version>>16) & 0xffff, int(d.version>>8) & 0xff, int(d.version) & 0xff
}

// eventTypes retrieves the specified event type, and passes it to f.
func (d *Evdev) eventTypes(typ, max int, f func(int)) error {
	buf := make([]uint64, max/64+1*(max%64))
	err := ioctl(d.fd.Fd(), _EVIOCGBIT(typ, max), unsafe.Pointer(&buf[0]))
	if err != nil {
		return err
	}
	for i := 0; i <= max; i++ {
		if (buf[i/64]>>uint(i%64))&1 == 1 {
			f(i)
		}
	}
	return nil
}

// EventTypes returns the device's supported event types.
func (d *Evdev) EventTypes() map[EventType]bool {
	var once sync.Once
	once.Do(func() {
		d.events = make(map[EventType]bool)
		d.eventTypes(int(EventSync), eventMax, func(i int) {
			d.events[EventType(i)] = true
		})
	})
	return d.events
}

// SyncTypes returns the sync events supported by the device.
//
// This is only applicable to devices with EventSync event support.
func (d *Evdev) SyncTypes() map[SyncType]bool {
	var once sync.Once
	once.Do(func() {
		d.syncs = make(map[SyncType]bool)
		d.eventTypes(int(EventSync), int(syncMax), func(i int) {
			d.syncs[SyncType(i)] = true
		})
	})
	return d.syncs
}

// KeyTypes returns the key events supported by the device.
//
// This is only applicable to devices with EventKey event support.
func (d *Evdev) KeyTypes() map[KeyType]bool {
	var once sync.Once
	once.Do(func() {
		d.keys = make(map[KeyType]bool)
		d.eventTypes(int(EventKey), int(keyMax), func(i int) {
			d.keys[KeyType(i)] = true
		})
	})
	return d.keys
}

// RelativeTypes returns a map of the supported relative axis types.
//
// This is only applicable to devices with EventRelative event support.
func (d *Evdev) RelativeTypes() map[RelativeType]bool {
	var once sync.Once
	once.Do(func() {
		d.relatives = make(map[RelativeType]bool)
		d.eventTypes(int(EventRelative), relativeMax, func(i int) {
			d.relatives[RelativeType(i)] = true
		})
	})
	return d.relatives
}

// AbsoluteTypes returns a map of the supported absolute axis types.
//
// This is only applicable to devices with EventAbsolute event support.
func (d *Evdev) AbsoluteTypes() map[AbsoluteType]Axis {
	var once sync.Once
	once.Do(func() {
		d.absolutes = make(map[AbsoluteType]Axis)
		d.eventTypes(int(EventAbsolute), absoluteMax, func(i int) {
			typ := AbsoluteType(i)
			d.absolutes[typ] = d.absoluteAxis(typ)
		})
	})
	return d.absolutes
}

// MiscTypes returns the misc events supported by the device.
//
// This is only applicable to devices with EventMisc event support.
func (d *Evdev) MiscTypes() map[MiscType]bool {
	var once sync.Once
	once.Do(func() {
		d.miscs = make(map[MiscType]bool)
		d.eventTypes(int(EventMisc), int(miscMax), func(i int) {
			d.miscs[MiscType(i)] = true
		})
	})
	return d.miscs
}

// SwitchTypes returns the switch events supported by the device.
//
// This is only applicable to devices with EventSwitch event support.
func (d *Evdev) SwitchTypes() map[SwitchType]bool {
	var once sync.Once
	once.Do(func() {
		d.switches = make(map[SwitchType]bool)
		d.eventTypes(int(EventSwitch), int(switchMax), func(i int) {
			d.switches[SwitchType(i)] = true
		})
	})
	return d.switches
}

// LEDTypes returns the led events supported by the device.
//
// This is only applicable to devices with EventLED event support.
func (d *Evdev) LEDTypes() map[LEDType]bool {
	var once sync.Once
	once.Do(func() {
		d.leds = make(map[LEDType]bool)
		d.eventTypes(int(EventLED), int(ledMax), func(i int) {
			d.leds[LEDType(i)] = true
		})
	})
	return d.leds
}

// SoundTypes returns the sound events supported by the device.
//
// This is only applicable to devices with EventSound event support.
func (d *Evdev) SoundTypes() map[SoundType]bool {
	var once sync.Once
	once.Do(func() {
		d.sounds = make(map[SoundType]bool)
		d.eventTypes(int(EventSound), int(soundMax), func(i int) {
			d.sounds[SoundType(i)] = true
		})
	})
	return d.sounds
}

// EffectTypes returns the force feedback effects supported by the
// device.
//
// This is only applicable to devices with EventEffect event support.
func (d *Evdev) EffectTypes() map[EffectType]bool {
	var once sync.Once
	once.Do(func() {
		d.effects = make(map[EffectType]bool)
		d.eventTypes(int(EventEffect), int(effectMax), func(i int) {
			d.effects[EffectType(i)] = true
		})
	})
	return d.effects
}

// PowerTypes returns the power events supported by the device.
//
// This is only applicable to devices with EventPower event support.
func (d *Evdev) PowerTypes() map[PowerType]bool {
	var once sync.Once
	once.Do(func() {
		d.powers = make(map[PowerType]bool)
		d.eventTypes(int(EventPower), int(powerMax), func(i int) {
			d.powers[PowerType(i)] = true
		})
	})
	return d.powers
}

// EffectStatusTypes returns the effects events supported by the device.
//
// This is only applicable to devices with EventEffectStatus event support.
func (d *Evdev) EffectStatusTypes() map[EffectStatusType]bool {
	var once sync.Once
	once.Do(func() {
		d.effectss = make(map[EffectStatusType]bool)
		d.eventTypes(int(EventEffectStatus), int(effectStatusMax), func(i int) {
			d.effectss[EffectStatusType(i)] = true
		})
	})
	return d.effectss
}

// EffectMax retrieves the maximum number of force feedback effects supported
// by the device.
//
// This is only applicable to devices with EventForceFeedback event support.
func (d *Evdev) EffectMax() int {
	var once sync.Once
	once.Do(func() {
		ioctl(d.fd.Fd(), _EVIOCGEFFECTS, unsafe.Pointer(&d.effectMax))
	})
	return int(d.effectMax)
}

// IsKeyboard returns true if the device qualifies as a keyboard.
func (d *Evdev) IsKeyboard() bool {
	m := d.EventTypes()
	return m[EventKey] && m[EventLED]
}

// IsMouse returns true if the device qualifies as a mouse.
func (d *Evdev) IsMouse() bool {
	m := d.EventTypes()
	return m[EventKey] && m[EventRelative]
}

// IsJoystick returns true if the device qualifies as a joystick.
func (d *Evdev) IsJoystick() bool {
	m := d.EventTypes()
	return m[EventKey] && m[EventAbsolute]
}

// absoluteAxis retrieves the state of the axis.
//
// If you want the global state for a device, you have to call the function for
// each axis present on the device.
//
// This is only applicable to devices with EventAbsolute event support.
func (d *Evdev) absoluteAxis(axis AbsoluteType) Axis {
	var abs Axis
	ioctl(d.fd.Fd(), _EVIOCGABS(int(axis)), unsafe.Pointer(&abs))
	return abs
}

// RepeatState returns the current, global repeat state. This applies only to
// devices which have the EventRepeat capability defined. This can be determined
// through `Device.EventTypes()`.
//
// Refer to Device.SetRepeatState for an explanation on what the returned
// values mean.
//
// This is only applicable to devices with EventRepeat event support.
func (d *Evdev) RepeatState() (uint, uint) {
	var rep [2]int32
	ioctl(d.fd.Fd(), _EVIOCGREP, unsafe.Pointer(&rep[0]))
	return uint(rep[0]), uint(rep[1])
}

// RepeatStateSet sets the repeat state for the given device.
//
// The values indicate (in milliseconds) the delay before the device starts
// repeating and the delay between subsequent repeats. This might apply to a
// keyboard where the user presses and holds a key.
//
// E.g.: We see an initial character immediately, then another @initial
// milliseconds later and after that, once every @subsequent milliseconds,
// until the key is released.
//
// This returns false if the operation failed.
//
// This is only applicable to devices with EventRepeat event support.
func (d *Evdev) RepeatStateSet(initial, subsequent uint) bool {
	rep := [2]int32{int32(initial), int32(subsequent)}
	return ioctl(d.fd.Fd(), _EVIOCSREP, unsafe.Pointer(&rep[0])) == nil
}

// KeyState returns the current, global key- and button- states.
//
// This is only applicable to devices with EventKey event support.
/*func (d *Evdev) KeyState() Bitset {
	b := NewBitset(keyMax)
	buf := b.Bytes()
	ioctl(d.fd.Fd(), _EVIOCGKEY(len(buf)), unsafe.Pointer(&buf[0]))
	return b
}*/

// KeyMap retrieves the key mapping for the given key.
func (d *Evdev) KeyMap(key KeyType) KeyMap {
	m := KeyMap{Key: uint32(key)}
	ioctl(d.fd.Fd(), _EVIOCGKEYCODE, unsafe.Pointer(&m))
	return m
}

// KeyMapSet sets a key map.
//
// This allows us to rewire physical keys -- ie, pressing M, will input N into
// the input system.
//
// Some input drivers support variable mappings between the keys held down
// (which are interpreted by the keyboard scan and reported as scancodes) and
// the events sent to the input layer.
//
// You can change which key is associated with each scancode using this call.
// The value of the scancode is the first element in the integer array
// (list[n][0]), and the resulting input event key number (keycode) is the
// second element in the array.  (list[n][1]).
//
// Be aware that the KeyMap functions may not work on every keyboard. This is
// only applicable to devices with EventKey event support.
func (d *Evdev) KeyMapSet(m KeyMap) error {
	return ioctl(d.fd.Fd(), _EVIOCSKEYCODE, unsafe.Pointer(&m))
}

// Poll polls the device for incoming events.
//
// Change the buffer size by specifying PollSize.
//
// Polling continues to run until the context is closed.
func (d *Evdev) Poll(ctxt context.Context) <-chan *EventEnvelope {
	count := d.pollSize
	if count == 0 {
		count = DefaultPollSize
	}

	ch := make(chan *EventEnvelope)
	go func() {
		defer close(ch)

		buf := make([]byte, sizeof_event*count)
		for {
			// check context
			select {
			case <-ctxt.Done():
				return
			default:
			}

			// read events
			i, err := d.fd.Read(buf)
			if err != nil {
				return
			}
			events := (*(*[1<<27 - 1]Event)(unsafe.Pointer(&buf[0])))[:i/sizeof_event]
			for _, e := range events {
				switch e.Type {
				case EventSync:
					ch <- &EventEnvelope{e, SyncType(e.Code)}
				case EventKey:
					ch <- &EventEnvelope{e, KeyType(e.Code)}
				case EventRelative:
					ch <- &EventEnvelope{e, RelativeType(e.Code)}
				case EventAbsolute:
					ch <- &EventEnvelope{e, AbsoluteType(e.Code)}
				case EventMisc:
					ch <- &EventEnvelope{e, MiscType(e.Code)}
				case EventSwitch:
					ch <- &EventEnvelope{e, SwitchType(e.Code)}
				case EventLED:
					ch <- &EventEnvelope{e, LEDType(e.Code)}
				case EventSound:
					ch <- &EventEnvelope{e, SoundType(e.Code)}
				case EventRepeat:
					ch <- &EventEnvelope{e, RepeatType(e.Code)}
				case EventEffect:
					ch <- &EventEnvelope{e, EffectType(e.Code)}
				case EventPower:
					ch <- &EventEnvelope{e, PowerType(e.Code)}
				case EventEffectStatus:
					ch <- &EventEnvelope{e, EffectStatusType(e.Code)}
				default:
					ch <- &EventEnvelope{e, nil}
				}
			}
		}
	}()
	return ch
}

// Send sends an event to the device.
func (d *Evdev) Send(ev Event) error {
	var once sync.Once
	once.Do(func() {
		var ctxt context.Context
		ctxt, d.cancel = context.WithCancel(context.Background())
		d.out = make(chan Event, 1)
		go func() {
			defer close(d.out)
			var event Event
			for {
				select {
				case event = <-d.out:
					buf := (*(*[1<<27 - 1]byte)(unsafe.Pointer(&event)))[:sizeof_event]
					n, err := d.fd.Write(buf)
					if err != nil {
						return
					}
					if n < sizeof_event {
						fmt.Fprintf(os.Stderr, "poll outbox: short write\n")
					}

				case <-ctxt.Done():
					break
				}
			}
		}()
	})
	d.out <- ev
	return nil
}

// EffectSet sends the Force Feedback effect to the device. The number of
// effects sent should not exceed the length of the device's EffectTypes.
//
// After this call completes, the effect.ID field will contain the effect's ID.
// The effect's ID must be used when playing or stopping the effect. It is also
// possible to reupload the same effect with the same ID later on with new
// parameters.
//
// This allows us to update a running effect, without first stopping it.
//
// This is only applicable to devices with EventForceFeedback event support.
func (d *Evdev) EffectSet(effect *Effect) error {
	return ioctl(d.fd.Fd(), _EVIOCSFF, unsafe.Pointer(effect))
}

// EffectUnset deletes the given effects from the device. This makes room for
// new effects in the device's memory. Note that this also stops the effect if
// it was playing.
//
// This is only applicable to devices with EventForceFeedback event support.
func (d *Evdev) EffectUnset(effect *Effect) error {
	return ioctl(d.fd.Fd(), _EVIOCRMFF, int(effect.ID))
}

// effectSend sends the specified effect with the value.
func (d *Evdev) effectSend(id EffectType, value int32) {
	d.Send(Event{
		Type:  EventEffect,
		Code:  uint16(id),
		Value: value,
	})
}

// EffectPlay plays a previously uploaded effect.
func (d *Evdev) EffectPlay(id EffectType) {
	d.effectSend(id, 1)
}

// EffectStop stops a previously uploaded effect from playing.
func (d *Evdev) EffectStop(id EffectType) {
	d.effectSend(id, 0)
}

// effectSet changes the given effect factor.
func (d *Evdev) effectPropSet(code EffectPropType, factor int) {
	if factor < 0 {
		factor = 0
	}
	if factor > 100 {
		factor = 100
	}
	d.Send(Event{
		Type:  EventEffect,
		Code:  uint16(code),
		Value: 0xffff * int32(factor) / 100,
	})
}

// EffectGainSet changes the force feedback gain.
//
// Not all devices have the same effect strength. Therefore, users should set a
// gain factor depending on how strong they want effects to be. This setting is
// persistent across access to the driver.
//
// The specified gain should be in the range 0-100. This is only applicable to
// devices with EventForceFeedback event support.
func (d *Evdev) EffectGainSet(gain int) {
	d.effectPropSet(EffectPropGain, gain)
}

// EffectAutoCenterSet changes the force feedback autocenter factor.
// The specified factor should be in the range 0-100.
// A value of 0 means: no autocenter.
//
// This is only applicable to devices with EventForceFeedback event support.
func (d *Evdev) EffectAutoCenterSet(factor int) {
	d.effectPropSet(EffectPropAutoCenter, factor)
}
