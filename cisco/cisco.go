// Package cisco implements Parser interface from github.com/doctor-fate/mskix
package cisco

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"

	"github.com/doctor-fate/mskix"
	"github.com/doctor-fate/mskix-parsers/utility"
	"github.com/doctor-fate/mskix/device"
)

// Id is the name of the current parser. Must be unique
// See github.com/doctor-fate/mskix-parsers for all available parsers
const Id device.ID = "Cisco"

type parser struct {
	header *regexp.Regexp
	re     *regexp.Regexp
}

func init() {
	mskix.Register(Id, newParser())
}

func newParser() mskix.Parser {
	return &parser{
		header: regexp.MustCompile(`\s*Port\s+Name\s+Status\s+Vlan\s+Duplex\s+Speed\s+Type\s+`),
		re: regexp.MustCompile(
			`\s*([[:alpha:]]{2}\d+/\d+)\s+([[:graph:]]*)\s+(?:disabled|connected)\s+(\d+)\s+(?:auto|full)\s+(?:auto|\d+)\s+(?:[[:alnum:]]*)\s*`),
	}
}

// Parse implements github.com/doctor-fate/mskix.Parser
// It reads input line by line and parses each line using parseRecord function
// Returns device.Data and error, if any. Error always nil, except header doesn't match
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

// parseRecord parses one line of input using regex and returns device.Record and error, if no match was found
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
	v := s != ""
	return device.NewEmptyString(s, v)
}
