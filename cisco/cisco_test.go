package cisco

import (
	"github.com/doctor-fate/mskix/device"
	"testing"
)

func TestParser_Parse_Valid(t *testing.T) {
	p := newParser()

	data, err := p.Parse([]byte(validInput))
	if err != nil {
		t.Errorf("err: expected nil, got %s", err)
	}
	if n := len(data.Records); n != 6 {
		t.Errorf("len(data.Records): expected 6, got %d", n)
	}

	r3 := device.Record{
		Port:        "Fa0/4",
		Description: device.NewEmptyString("dwecwf==03233-08", true),
		VLAN:        device.NewEmptyString("99", true),
	}
	if r := data.Records[3]; r != r3 {
		t.Errorf("data.Records[3]: expected %v, got %v", r3, r)
	}

	r0 := device.Record{
		Port:        "Fa0/1",
		Description: device.NewEmptyString("", false),
		VLAN:        device.NewEmptyString("99", true),
	}
	if r := data.Records[0]; r != r0 {
		t.Errorf("data.Records[0]: expected %v, got %v", r0, r)
	}
}

func TestParser_Parse_Invalid(t *testing.T) {
	p := newParser()
	for _, v := range []string{invalidInput, invalidInput1, invalidInput2} {
		_, err := p.Parse([]byte(v))
		if err == nil {
			t.Errorf("err: expected not nil err, got nil")
		}
	}
}

var validInput = ` Port      Name               Status       Vlan       Duplex Speed Type
 Fa0/1                        disabled     99           auto auto 10/100BaseTX
 Fa0/2                        disabled     99           auto auto 10/100BaseTX
 Fa0/3        ffftk                disabled     99           auto auto 10/100BaseTX
 Fa0/4          dwecwf==03233-08              disabled     99           auto auto 10/100BaseTX
 Fa0/5                        disabled     99           auto auto 10/100BaseTX
 Fa0/6       fwefewfew                 disabled     99           auto auto 10/100BaseTX`

var invalidInput = ` Port      Name               Status1       Vlan       Duplex Speed Type
 Fa0/1                        disabled     99           auto auto 10/100BaseTX
 Fa0/2                        disabled     99           auto auto 10/100BaseTX
 Fa0/3                        disabled     99           auto auto 10/100BaseTX
 Fa0/4                        disabled     99           auto auto 10/100BaseTX
 Fa0/5                        disabled     99           auto auto 10/100BaseTX
 Fa0/6                        disabled     99           auto auto 10/100BaseTX`

var invalidInput1 = ` Port                 Status       Vlan       Duplex Speed Type
 Fa0/1                        disable     99           auto auto 10/100BaseTX
 Fa0/2                        disabled     99           auto auto 10/100BaseTX
 Fa0/3                        disabled     99           auto auto 10/100BaseTX
 Fa0/4                        disabled     99           auto auto 10/100BaseTX
 Fa0/5                        disabled     99           auto auto 10/100BaseTX
 Fa0/6                        disabled     99           auto auto 10/100BaseTX`

var invalidInput2 = ` Port      Name               Status       Vlan       Duplex Speed Typed
 Fa0/1                        disabled     99           auto auto 10/100BaseTX
 Fa0/2                        disabled     99           auto auto 10/100BaseTX
 Fa0/3                        disabled                auto auto 10/100BaseTX
 Fa0/4                        disabled     99           auto auto 10/100BaseTX
 Fa0/5                        disabled     99           auto auto 10/100BaseTX
 Fa0/6                        disabled     99           auto auto 10/100BaseTX`
