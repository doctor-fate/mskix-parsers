package extreme

import (
	"testing"

	"github.com/doctor-fate/mskix/device"
)

func TestParser_Parse_Valid(t *testing.T) {
	p := newParser()

	data, err := p.Parse([]byte(validInput))
	if err != nil {
		t.Errorf("err: expected nil, got %s", err)
	}
	if n := len(data.Records); n != 13 {
		t.Errorf("len(data.Records): expected 13, got %d", n)
	}

	r3 := device.Record{
		Port:        "4",
		Description: device.NewEmptyString("", false),
		VLAN:        device.NewEmptyString("VLAN0099", true),
	}
	if r := data.Records[3]; r != r3 {
		t.Errorf("data.Records[3]: expected %v, got %v", r3, r)
	}

	r0 := device.Record{
		Port:        "1",
		Description: device.NewEmptyString("mfff-10.3", true),
		VLAN:        device.NewEmptyString("(0005)", true),
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

var validInput = `Port Summary
  Port  Display              VLAN Name          Port  Link  Speed Duplex
  #     String               (or # VLANs)       State State Actual Actual
  =======================================================================
  1     mfff-10.3            (0005)              E     A     10G FULL
  2     fds-35.3             VLAN0901            D     R
  3     yyyyyx-46            VLAN0901            E     A     10G FULL
  4                          VLAN0099            D     R
  5     mmfsd-10.2           (0005)              E     A     10G FULL
  6     sfdsfsf-8            VLAN0901            E     A     10G FULL
  7     fsd-35.2             VLAN0901            D     R
  8     ggggaa-12.2          VLAN0901            E     A     10G FULL
  9     ggggaa-12.3          VLAN0901            E     A     10G FULL
  10    vfsfssfe-4.2         VLAN0901            E     A     10G FULL
  11    vfsfssfe-4.3         VLAN0901            E     A     10G FULL
  12    mfff-9               VLAN0962            E     A     10G FULL
  13    nfsfsfs-5            VLAN0901            E     A     10G FULL`

var invalidInput = `Port Summary
  Port  Display              VLAN Name          Port  Link Column Speed Duplex
  #     String               (or # VLANs)       State State Actual Actual
  =======================================================================
  1     mfff-10.3            (0005)              E     A     10G FULL
  2     fds-35.3             VLAN0901            D     R
  3     yyyyyx-46            VLAN0901            E     A     10G FULL
  4                          VLAN0099            D     R
  5     mmfsd-10.2           (0005)              E     A     10G FULL
  6     sfdsfsf-8            VLAN0901            E     A     10G FULL
  7     fsd-35.2             VLAN0901            D     R
  8     ggggaa-12.2          VLAN0901            E     A     10G FULL
  9     ggggaa-12.3          VLAN0901            E     A     10G FULL
  10    vfsfssfe-4.2         VLAN0901            E     A     10G FULL
  11    vfsfssfe-4.3         VLAN0901            E     A     10G FULL
  12    mfff-9               VLAN0962            E     A     10G FULL
  13    nfsfsfs-5            VLAN0901            E     A     10G FULL`

var invalidInput1 = `Port Summary
  Port  Display              VLANName          Port  Link  Speed Duplex
  #     String               (or # VLANs)       State State Actual Actual
  =======================================================================
  1     mfff-10.3            (0005)              E     A     10G FULL
  2     fds-35.3             VLAN0901            D     R
  3     yyyyyx-46            VLAN0901                 A     10G FULL
  4                          VLAN0099            D     R
  5     mmfsd-10.2           (0005)              E     A     10G FULL
  6     sfdsfsf-8            VLAN0901            E     A     10G FULL
  7     fsd-35.2             VLAN0901            D     R
  8     ggggaa-12.2          VLAN0901            E     A     10G FULL
  9     ggggaa-12.3          VLAN0901            E     A     10G FULL
  10    vfsfssfe-4.2         VLAN0901            E     A     10G FULL
  11    vfsfssfe-4.3         VLAN0901            E     A     10G FULL
  12    mfff-9               VLAN0962            E     A     10G FULL
  13    nfsfsfs-5            VLAN0901            E     A     10G FULL`

var invalidInput2 = `Port Summary
  Port  Display              VLAN Name         Link  Speed Duplex
  #     String               (or # VLANs)       State State Actual Actual
  =======================================================================
  1     mfff-10.3            (0005)              E     A     10G FULL
  2     fds-35.3             VLAN0901            D     R
  3     yyyyyx-46            VLAN0901            E     A     10G FULL
  4                          VLAN0099            D     R
  5     mmfsd-10.2           (0005)              E     A     10G FULL
  6     sfdsfsf-8            VLAN0901            E     A     10G FULL
  7     fsd-35.2             VLAN0901            D     R
  8     ggggaa-12.2          VLAN0901            E     A     10G FULL
  9     ggggaa-12.3          VLAN0901            E     A     10G FULL
  10    vfsfssfe-4.2         VLAN0901            E     A     10G FULL
  11    vfsfssfe-4.3         VLAN0901            E     A     10G FULL
  12    mfff-9               VLAN0962            E     A     10G FULL
  13    nfsfsfs-5            VLAN0901            E     A     10G FULL`
