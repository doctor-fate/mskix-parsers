package force10

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"

	"github.com/doctor-fate/mskix"
	"github.com/doctor-fate/mskix-parsers/utility"
	"github.com/doctor-fate/mskix/device"
)

const Id device.ID = "Force10"

type parser struct {
	header *regexp.Regexp
	re     *regexp.Regexp
}

func init() {
	mskix.Register(Id, newParser())
}

func newParser() mskix.Parser {
	return &parser{
		header: regexp.MustCompile(`\s*Port\s+Description\s+Status\s+Speed\s+Duplex\s+Vlan\s+`),
		re: regexp.MustCompile(
			`\s*([[:alpha:]]{2} \d+/\d+)\s+([[:graph:]]*)\s+(?:Up|Down)\s+(?:Auto|\d+ Mbit)\s+(?:Full|Auto|Half)\s+([[:graph:]]*)\s*`),
	}
}

func (p *parser) Parse(input []byte) (device.Data, error) {
	var data = device.Data{
		Id: Id,
	}
	r := bufio.NewReader(bytes.NewReader(input))
	header, err := r.ReadBytes('\n')
	if err != nil {
		return data, err
	}
	if !p.header.Match(header) {
		return data, fmt.Errorf("Parse: header doesn't match: id=%s, header=%s", Id, header)
	}

	data.Records, err = utility.ReadInput(r, p.parseRecord)
	return data, err
}

func (p *parser) parseRecord(input []byte) (device.Record, error) {
	matches := p.re.FindSubmatch(input)
	if matches == nil {
		return device.Record{}, fmt.Errorf("parseRecord: invalid record: id=%s, input=%s", Id, input)
	}
	return device.Record{
		Port:        string(matches[1]),
		Description: p.sanitize(matches[2]),
		VLAN:        p.sanitize(matches[3]),
	}, nil
}

func (p *parser) sanitize(input []byte) device.EmptyString {
	s := string(input)
	v := s != "" && s != "--"
	return device.NewEmptyString(s, v)
}
