# GPIO

This package provides drivers for [General Purpose Input/Output (GPIO)](https://en.wikipedia.org/wiki/General_Purpose_Input/Output) devices. It is normally used by connecting an adaptor such as [firmata](https://github.com/eyelight/gobot/platforms/firmata) that supports the needed interfaces for GPIO devices.

## Getting Started

## Installing
```
go get -d -u github.com/eyelight/gobot/...
```

## Hardware Support
Gobot has a extensible system for connecting to hardware devices. The following GPIO devices are currently supported:
	- Button
	- Buzzer
	- Direct Pin
	- Grove Button
	- Grove Buzzer
	- Grove LED
	- Grove Magnetic Switch
	- Grove Relay
	- Grove Touch Sensor
	- LED
	- Makey Button
	- Motor
	- Proximity Infra Red (PIR) Motion Sensor
	- Relay
	- RGB LED
	- Servo
	- Stepper Motor
	- TM1638 LED Controller

More drivers are coming soon...
