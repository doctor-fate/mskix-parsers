package force10

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
	if n := len(data.Records); n != 16 {
		t.Errorf("len(data.Records): expected 15, got %d", n)
	}

	r11 := device.Record{
		Port:        "Gi 0/11",
		Description: device.NewEmptyString("ttttttttt-5", true),
		VLAN:        device.NewEmptyString("--", false),
	}
	if r := data.Records[11]; r != r11 {
		t.Errorf("data.Records[3]: expected %v, got %v", r11, r)
	}

	r3 := device.Record{
		Port:        "Gi 0/3",
		Description: device.NewEmptyString("", false),
		VLAN:        device.NewEmptyString("99", true),
	}
	if r := data.Records[3]; r != r3 {
		t.Errorf("data.Records[0]: expected %v, got %v", r3, r)
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

var validInput = `Port     Description  Status Speed     Duplex Vlan
Gi 0/0   abcdd-10.2   Up     100 Mbit  Full   786-787,999
Gi 0/1   defgg        Up     100 Mbit  Full   828
Gi 0/2   rfv-4.2      Up     1000 Mbit Full   743
Gi 0/3                Down   Auto      Auto   99
Gi 0/4   frgts-9      Up     100 Mbit  Full   701
Gi 0/5   adfdfd-2     Up     1000 Mbit Full   701
Gi 0/6   yuuig.2      Up     100 Mbit  Full   798
Gi 0/7   rrrrrr.2     Up     1000 Mbit Full   858
Gi 0/8   ttttttttt-5  Up     1000 Mbit Full   --
Gi 0/9                Down   Auto      Auto   99
Gi 0/10               Down   Auto      Auto   99
Gi 0/11  ttttttttt-5  Up     1000 Mbit Full   --
Gi 0/12  dfsfs        Up     1000 Mbit Full   701
Gi 0/13               Down   Auto      Auto   99
Gi 0/14               Down   Auto      Auto   99
Gi 0/15  gggm-3       Up     10 Mbit   Full   406`

var invalidInput = `   Description  Status Speed     Duplex Vlan
Gi 0/0   abcdd-10.2   Up     100 Mbit  Full   786-787,999
Gi 0/1   defgg        Up     100 Mbit  Full   828
Gi 0/2   rfv-4.2      Up     1000 Mbit Full   743
Gi 0/3                Down   Auto      Auto   99
Gi 0/4   frgts-9      Up     100 Mbit  Full   701
Gi 0/5   adfdfd-2     Up     1000 Mbit Full   701
Gi 0/6   yuuig.2      Up     100 Mbit  Full   798
Gi 0/7   rrrrrr.2     Up     1000 Mbit Full   858
Gi 0/8   ttttttttt-5  Up     1000 Mbit Full   --
Gi 0/9                Down   Auto      Auto   99
Gi 0/10               Down   Auto      Auto   99
Gi 0/11  ttttttttt-5  Up     1000 Mbit Full   --
Gi 0/12  dfsfs        Up     1000 Mbit Full   701
Gi 0/13               Down   Auto      Auto   99
Gi 0/14               Down   Auto      Auto   99
Gi 0/15  gggm-3       Up     10 Mbit   Full   406`

var invalidInput1 = `Port Du    Description  Status Speed     Duplex Vlan
Gi 0/0   abcdd-10.2   Up     100 Mbit  Full   786-787,999
Gi 0/1   defgg        Up     100 Mbit  Full   828
Gi 0/2   rfv-4.2      Up     1000 Mbit Full   743
Gi 0/3                Down   Auto      Auto   99
Gi 0/4   frgts-9      Up     100 Mbit  Full   701
Gi 0/5   adfdfd-2     Up     1000 Mbit Full   701
Gi 0/6   yuuig.2      Up     100 Mbit  Full   798
Gi 0/7   rrrrrr.2     Up     1000 Mbit Full   858
Gi 0/8   ttttttttt-5  Up     1000 Mbit Full   --
Gi 0/9                Down   Auto      Auto   99
Gi 0/10               Down   Auto      Auto   99
Gi 0/11  ttttttttt-5  Up     1000 Mbit Full   --
Gi 0/12  dfsfs        Up     1000 Mbit Full   701
Gi 0/13               Down   Auto      Auto   99
Gi 0/14               Down   Auto      Auto   99
Gi 0/15  gggm-3       Up     10 Mbit   Full   406`

var invalidInput2 = `Gi 0/0   abcdd-10.2   Up     100 Mbit  Full   786-787,999
Gi 0/1   defgg        Up     100 Mbit  Full   828
Gi 0/2   rfv-4.2      Up     1000 Mbit Full   743
Gi 0/3                Down   Auto      Auto   99
Gi 0/4   frgts-9      Up     100 Mbit  Full   701
Gi 0/5   adfdfd-2     Up     1000 Mbit Full   701
Gi 0/6   yuuig.2      Up     100 Mbit  Full   798
Gi 0/7   rrrrrr.2     Up     1000 Mbit Full   858
Gi 0/8   ttttttttt-5  Up     1000 Mbit Full   --
Gi 0/9                Down   Auto      Auto   99
Gi 0/10               Down   Auto      Auto   99
Gi 0/11  ttttttttt-5  Up     1000 Mbit Full   --
Gi 0/12  dfsfs        Up     1000 Mbit Full   701
Gi 0/13               Down   Auto      Auto   99
Gi 0/14               Down   Auto      Auto   99
Gi 0/15  gggm-3       Up     10 Mbit   Full   406`
