# About evdev examples

This directory contains project examples for using the [`evdev`][1] package.

## Available examples

The following examples are currently available:

<!-- the following section is updated by running `go run gen.go` -->
<!-- START EXAMPLES -->
| Example           | Description                                                                       |
|-------------------|-----------------------------------------------------------------------------------|
| [ff](/ff)         | send force feedback events to a input device                                      |
| [iter](/iter)     | iterate input devices and read their configuration                                |
| [led](/led)       | send events (LED) to an input device                                              |
| [poll](/poll)     | poll a input device for events                                                    |
| [relay](/relay)   | relaying events from an input device to a uinput device                           |
| [uinput](/uinput) | create a virtual user input ("uinput") device and sending/receiving events for it |
<!-- END EXAMPLES -->

[1]: https://github.com/kenshaw/evdev
