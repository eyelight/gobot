package bb8

import (
	"strings"
	"testing"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/gobottest"
)

var _ gobot.Driver = (*BB8Driver)(nil)

func initTestBB8Driver() *BB8Driver {
	d := NewDriver(NewBleTestAdaptor())
	return d
}

func TestBB8Driver(t *testing.T) {
	d := initTestBB8Driver()
	gobottest.Assert(t, strings.HasPrefix(d.Name(), "BB8"), true)
	d.SetName("NewName")
	gobottest.Assert(t, d.Name(), "NewName")
}

func TestBB8DriverStartAndHalt(t *testing.T) {
	d := initTestBB8Driver()
	gobottest.Assert(t, d.Start(), nil)
	gobottest.Assert(t, d.Halt(), nil)
}
