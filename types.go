package evdev

import (
	"syscall"
	"unsafe"
)

// ID represents the device identity.
//
// The bus type is the only field that contains accurate data. It can be
// compared to the BusXXX constants.
//
// The vendor, product and version fields are bus type-specific information
// relating to the identity of the device.
//
// Modern devices (typically using PCI or USB) do have information that can be
// used, but legacy devices (such as serial mice,
//
// PS/2 keyboards and game ports on ISA sound cards) do not. These numbers
// therefore are not meaningful for some values of bus type.
type ID struct {
	BusType BusType
	Vendor  uint16
	Product uint16
	Version uint16
}

// Axis provides information for a specific absolute axis. This applies to
// devices which support EventAbsolute events.
type Axis struct {
	// Val is the current value of the axis.
	Val int32

	// Min is the axis minimum.
	Min int32 // Lower limit of axis.

	// Max is the axis maximum.
	Max int32 // Upper limit of axis.

	// Fuzz is the axis fuzz factor.
	Fuzz int32 // ???

	// Flat is the size of the axis' flat section.
	Flat int32

	// Res is the size of any error present.
	Res int32
}

// Event represents a generic input event.
type Event struct {
	Time  syscall.Timeval
	Type  EventType
	Code  uint16
	Value int32
}

// KeyMap is used to retrieve and modify keymap data. Users have option of
// performing lookup either by @scancode itself or by @index in a keymap entry.
// Device.KeyMap() will also return scancode or index (depending on which
// element was used to perform lookup).
type KeyMap struct {
	Flags uint8     // They specify how the kernel should handle a keymap request.
	Len   uint8     // Length of the scancode that resides in Scancode buffer.
	Index uint16    // Index in the keymap, may be used instead of scancode
	Key   uint32    // Key code assigned to this scancode
	Code  [32]uint8 // Scancode represented in machine-endian form.
}

// Effect describes a force feedback effect.
//
// Supported effects are as follows:
//
// 	- EffectRumble          rumble effects
// 	- EffectPeriodic        periodic effects (using waveform):
// 	  - EffectSquare        square waveform
// 	  - EffectTriangle      triangle waveform
// 	  - EffectSine          sine waveform
// 	  - EffectSawUp         sawtooth up waveform
// 	  - EffectSawDown       sawtooth down waveform
// 	  - EffectCustom        custom waveform
// 	- EffectConstant        renders constant force effects
// 	- EffectSpring          simulates the presence of a spring
// 	- EffectFriction        simulates friction
// 	- EffectDamper          simulates damper effects
// 	- EffectInertia         simulates inertia
// 	- EffectRamp            simulates ramp effects
//
//  Supported adjustments:
//
// 	- EffectPropGain        gain adjust
// 	- EffectPropAutoCenter  auto center adjust
//
// Note: In most cases you should use EffectPeriodic instead of EffectRumble.
// All devices that support EffectRumble support EffectPeriodic (square,
// triangle, sine) and the other way around.
//
// Note: The exact layout of EffectCustom waveforms is undefined for the time
// being as no driver supports it yet.
//
// Note: All duration values are expressed in milliseconds. Values above 32767
// ms (0x7fff) should not be used and have unspecified results.
type Effect struct {
	Type      EffectType
	ID        int16
	Direction EffectDirType
	Trigger   EffectTrigger
	Replay    EffectReplay
	data      unsafe.Pointer
}

// EffectOption is an effect option.
type EffectOption func(*Effect)

// NewEffect wraps creating an effect.
func NewEffect(typ EffectType, opts ...EffectOption) *Effect {
	e := &Effect{
		Type: typ,
		ID:   -1,
	}
	for _, o := range opts {
		o(e)
	}
	return e
}

// EffectID is an effect option to set the effect's ID.
func EffectID(id int) EffectOption {
	return func(e *Effect) {
		e.ID = int16(id)
	}
}

// EffectDirection is an effect option to set the effect's direction.
func EffectDirection(dir EffectDirType) EffectOption {
	return func(e *Effect) {
		e.Direction = dir
	}
}

// Data returns the event data structure as a concrete type. Its type depends
// on the value of Effect.Type and can be any of:
//
//    EffectConstant -> ConstantEffect
//    EffectPeriodic -> PeriodicEffect
//    EffectRamp     -> RampEffect
//    EffectRumble   -> RumbleEffect
//    EffectSpring   -> [2]ConditionEffect
//    EffectDamper   -> [2]ConditionEffect
//
// This returns nil if the type was not recognized.
func (e *Effect) Data() interface{} {
	// FIXME(jimt): Deal with: EffectFriction, EffectInertia: Unsure what they
	// should return.
	if e.data == nil {
		return nil
	}

	switch e.Type {
	case EffectRumble:
		return *(*RumbleEffect)(e.data)
	case EffectPeriodic:
		return *(*PeriodicEffect)(e.data)
	case EffectConstant:
		return *(*ConstantEffect)(e.data)
	case EffectSpring, EffectDamper:
		return *(*[2]ConditionEffect)(e.data)
	case EffectRamp:
		return *(*RampEffect)(e.data)
	}

	return nil
}

// SetData sets the event data structure.
func (e *Effect) SetData(v interface{}) {
	if v != nil {
		e.data = unsafe.Pointer(&v)
	}
}

// EffectReplay ...
type EffectReplay struct {
	Length uint16
	Delay  uint16
}

// EffectTrigger ...
type EffectTrigger struct {
	Button   uint16
	Interval uint16
}

// EffectEnvelope ...
type EffectEnvelope struct {
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

// RumbleEffect wraps the data for a rumble force feedback effect.
//
// The rumble effect is the most basic effect, causing the device to vibrate.
// The API contains support for two motors, a strong one and a weak one, which
// can be controlled independently.
type RumbleEffect struct {
	MagnitudeStrong uint16
	MagnitudeWeak   uint16
}

// ConstantEffect wraps the data for a constant force feedback effect.
type ConstantEffect struct {
	Level    int16
	Envelope EffectEnvelope
}

// RampEffect wraps the data for a ramp force feedback effect.
type RampEffect struct {
	LevelStart int16
	LevelEnd   int16
	Envelope   EffectEnvelope
}

// ConditionEffect wraps the data for a condition force feedback effect.
type ConditionEffect struct {
	SaturationRight  uint16
	SaturationLeft   uint16
	CoeffecientRight int16
	CoeffecientLeft  int16
	Deadband         uint16
	Center           int16
}

// PeriodicEffect wraps the data for a periodic (waveform) force feedback
// effect.
//
// waveforms: Square, Triangle, Sine, Sawtooth or a custom waveform.
type PeriodicEffect struct {
	Waveform  uint16
	Period    uint16
	Magnitude int16
	Offset    int16
	Phase     uint16
	Envelope  EffectEnvelope

	custom_len  uint32
	custom_data unsafe.Pointer // *int16
}

// Data returns custom waveform information. This comes in the form of a signed
// 16-bit slice.
//
// The exact layout of a custom waveform is undefined for the time being as no
// driver supports it yet.
func (e *PeriodicEffect) Data() []int16 {
	if e.custom_data == nil {
		return nil
	}
	return (*(*[1<<27 - 1]int16)(e.custom_data))[:e.custom_len]
}

// SetData sets custom waveform information.
//
// The exact layout of a custom waveform is undefined for the time being as no
// driver supports it yet.
func (e *PeriodicEffect) SetData(v []int16) {
	e.custom_len = uint32(len(v))
	e.custom_data = unsafe.Pointer(nil)

	if e.custom_len > 0 {
		e.custom_data = unsafe.Pointer(&v[0])
	}
}
