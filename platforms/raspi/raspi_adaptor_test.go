package raspi

import (
	"errors"
	"strings"
	"testing"

	"runtime"
	"strconv"
	"sync"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/drivers/gpio"
	"github.com/eyelight/gobot/drivers/i2c"
	"github.com/eyelight/gobot/drivers/spi"
	"github.com/eyelight/gobot/gobottest"
	"github.com/eyelight/gobot/sysfs"
)

// make sure that this Adaptor fullfills all the required interfaces
var _ gobot.Adaptor = (*Adaptor)(nil)
var _ gpio.DigitalReader = (*Adaptor)(nil)
var _ gpio.DigitalWriter = (*Adaptor)(nil)
var _ gpio.PwmWriter = (*Adaptor)(nil)
var _ gpio.ServoWriter = (*Adaptor)(nil)
var _ sysfs.DigitalPinnerProvider = (*Adaptor)(nil)
var _ sysfs.PWMPinnerProvider = (*Adaptor)(nil)
var _ i2c.Connector = (*Adaptor)(nil)
var _ spi.Connector = (*Adaptor)(nil)

func initTestAdaptor() *Adaptor {
	readFile = func() ([]byte, error) {
		return []byte(`
Hardware        : BCM2708
Revision        : 0010
Serial          : 000000003bc748ea
`), nil
	}
	a := NewAdaptor()
	a.Connect()
	return a
}

func TestRaspiAdaptorName(t *testing.T) {
	a := initTestAdaptor()
	gobottest.Assert(t, strings.HasPrefix(a.Name(), "RaspberryPi"), true)
	a.SetName("NewName")
	gobottest.Assert(t, a.Name(), "NewName")
}

func TestAdaptor(t *testing.T) {
	readFile = func() ([]byte, error) {
		return []byte(`
Hardware        : BCM2708
Revision        : 0010
Serial          : 000000003bc748ea
`), nil
	}
	a := NewAdaptor()
	gobottest.Assert(t, strings.HasPrefix(a.Name(), "RaspberryPi"), true)
	gobottest.Assert(t, a.i2cDefaultBus, 1)
	gobottest.Assert(t, a.revision, "3")

	readFile = func() ([]byte, error) {
		return []byte(`
Hardware        : BCM2708
Revision        : 000D
Serial          : 000000003bc748ea
`), nil
	}
	a = NewAdaptor()
	gobottest.Assert(t, a.i2cDefaultBus, 1)
	gobottest.Assert(t, a.revision, "2")

	readFile = func() ([]byte, error) {
		return []byte(`
Hardware        : BCM2708
Revision        : 0002
Serial          : 000000003bc748ea
`), nil
	}
	a = NewAdaptor()
	gobottest.Assert(t, a.i2cDefaultBus, 0)
	gobottest.Assert(t, a.revision, "1")

}

func TestAdaptorFinalize(t *testing.T) {
	a := initTestAdaptor()

	fs := sysfs.NewMockFilesystem([]string{
		"/sys/class/gpio/export",
		"/sys/class/gpio/unexport",
		"/dev/pi-blaster",
		"/dev/i2c-1",
		"/dev/i2c-0",
		"/dev/spidev0.0",
		"/dev/spidev0.1",
	})

	sysfs.SetFilesystem(fs)
	sysfs.SetSyscall(&sysfs.MockSyscall{})

	a.DigitalWrite("3", 1)
	a.PwmWrite("7", 255)

	a.GetConnection(0xff, 0)
	gobottest.Assert(t, a.Finalize(), nil)
}

func TestAdaptorDigitalPWM(t *testing.T) {
	a := initTestAdaptor()
	a.PiBlasterPeriod = 20000000

	gobottest.Assert(t, a.PwmWrite("7", 4), nil)

	fs := sysfs.NewMockFilesystem([]string{
		"/dev/pi-blaster",
	})
	sysfs.SetFilesystem(fs)

	pin, _ := a.PWMPin("7")
	period, _ := pin.Period()
	gobottest.Assert(t, period, uint32(20000000))

	gobottest.Assert(t, a.PwmWrite("7", 255), nil)

	gobottest.Assert(t, strings.Split(fs.Files["/dev/pi-blaster"].Contents, "\n")[0], "4=1")

	gobottest.Assert(t, a.ServoWrite("11", 90), nil)

	gobottest.Assert(t, strings.Split(fs.Files["/dev/pi-blaster"].Contents, "\n")[0], "17=0.5")

	gobottest.Assert(t, a.PwmWrite("notexist", 1), errors.New("Not a valid pin"))
	gobottest.Assert(t, a.ServoWrite("notexist", 1), errors.New("Not a valid pin"))

	pin, _ = a.PWMPin("12")
	period, _ = pin.Period()
	gobottest.Assert(t, period, uint32(20000000))

	gobottest.Assert(t, pin.SetDutyCycle(1.5*1000*1000), nil)

	gobottest.Assert(t, strings.Split(fs.Files["/dev/pi-blaster"].Contents, "\n")[0], "18=0.075")
}

func TestAdaptorDigitalIO(t *testing.T) {
	a := initTestAdaptor()
	fs := sysfs.NewMockFilesystem([]string{
		"/sys/class/gpio/export",
		"/sys/class/gpio/unexport",
		"/sys/class/gpio/gpio4/value",
		"/sys/class/gpio/gpio4/direction",
		"/sys/class/gpio/gpio27/value",
		"/sys/class/gpio/gpio27/direction",
	})

	sysfs.SetFilesystem(fs)

	a.DigitalWrite("7", 1)
	gobottest.Assert(t, fs.Files["/sys/class/gpio/gpio4/value"].Contents, "1")

	a.DigitalWrite("13", 1)
	i, _ := a.DigitalRead("13")
	gobottest.Assert(t, i, 1)

	gobottest.Assert(t, a.DigitalWrite("notexist", 1), errors.New("Not a valid pin"))

	fs.WithReadError = true
	_, err := a.DigitalRead("13")
	gobottest.Assert(t, err, errors.New("read error"))

	fs.WithWriteError = true
	_, err = a.DigitalRead("7")
	gobottest.Assert(t, err, errors.New("write error"))
}

func TestAdaptorI2c(t *testing.T) {
	a := initTestAdaptor()
	fs := sysfs.NewMockFilesystem([]string{
		"/dev/i2c-1",
	})
	sysfs.SetFilesystem(fs)
	sysfs.SetSyscall(&sysfs.MockSyscall{})

	con, err := a.GetConnection(0xff, 1)
	gobottest.Assert(t, err, nil)

	con.Write([]byte{0x00, 0x01})
	data := []byte{42, 42}
	con.Read(data)
	gobottest.Assert(t, data, []byte{0x00, 0x01})

	_, err = a.GetConnection(0xff, 51)
	gobottest.Assert(t, err, errors.New("Bus number 51 out of range"))

	gobottest.Assert(t, a.GetDefaultBus(), 1)
}

func TestAdaptorSPI(t *testing.T) {
	a := initTestAdaptor()
	fs := sysfs.NewMockFilesystem([]string{
		"/dev/spidev0.1",
	})
	sysfs.SetFilesystem(fs)
	sysfs.SetSyscall(&sysfs.MockSyscall{})

	gobottest.Assert(t, a.GetSpiDefaultBus(), 0)
	gobottest.Assert(t, a.GetSpiDefaultChip(), 0)
	gobottest.Assert(t, a.GetSpiDefaultMode(), 0)
	gobottest.Assert(t, a.GetSpiDefaultMaxSpeed(), int64(500000))

	_, err := a.GetSpiConnection(10, 0, 0, 8, 500000)
	gobottest.Assert(t, err.Error(), "Bus number 10 out of range")

	// TODO: test tx/rx here...
}

func TestAdaptorDigitalPinConcurrency(t *testing.T) {

	oldProcs := runtime.GOMAXPROCS(0)
	runtime.GOMAXPROCS(8)

	for retry := 0; retry < 20; retry++ {

		a := initTestAdaptor()
		var wg sync.WaitGroup

		for i := 0; i < 20; i++ {
			wg.Add(1)
			pinAsString := strconv.Itoa(i)
			go func(pin string) {
				defer wg.Done()
				a.DigitalPin(pin, sysfs.IN)
			}(pinAsString)
		}

		wg.Wait()
	}

	runtime.GOMAXPROCS(oldProcs)

}

func TestAdaptorPWMPin(t *testing.T) {
	a := initTestAdaptor()

	gobottest.Assert(t, len(a.pwmPins), 0)

	firstSysPin, err := a.PWMPin("35")

	gobottest.Assert(t, err, nil)
	gobottest.Assert(t, len(a.pwmPins), 1)

	secondSysPin, err := a.PWMPin("35")

	gobottest.Assert(t, err, nil)
	gobottest.Assert(t, len(a.pwmPins), 1)
	gobottest.Assert(t, firstSysPin, secondSysPin)

	otherSysPin, err := a.PWMPin("36")

	gobottest.Assert(t, err, nil)
	gobottest.Assert(t, len(a.pwmPins), 2)
	gobottest.Refute(t, firstSysPin, otherSysPin)
}
