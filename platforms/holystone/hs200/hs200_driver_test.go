package hs200

import (
	"testing"

	"github.com/eyelight/gobot"
	"github.com/eyelight/gobot/gobottest"
)

var _ gobot.Driver = (*Driver)(nil)

func TestHS200Driver(t *testing.T) {
	d := NewDriver("127.0.0.1:8080", "127.0.0.1:9090")

	gobottest.Assert(t, d.tcpaddress, "127.0.0.1:8080")
	gobottest.Assert(t, d.udpaddress, "127.0.0.1:9090")
}
