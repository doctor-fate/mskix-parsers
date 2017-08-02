package extreme

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"

	"github.com/doctor-fate/mskix"
	"github.com/doctor-fate/mskix/device"
)

const Id device.ID = "Extreme"

type parser struct {
	header *regexp.Regexp
	re     *regexp.Regexp
}

func init() {
	mskix.Register(Id, newParser())
}

func newParser() mskix.Parser {
	return &parser{
		header: regexp.MustCompile(`\s*Port\s+Display\s+VLAN Name\s+Port\s+Link\s+Speed\s+Duplex\s+`),
		re: regexp.MustCompile(
			`\s*(\d+)\s+([[:graph:]]*)\s+(VLAN\d+|\(\d+\))?\s+(?:[A-Z])\s+(?:[A-Z]+)(?:\s+\d+G*\s+FULL)?\s+`),
	}
}

func (p *parser) Parse(input []byte) (device.Data, error) {
	var data = device.Data{
		Id: Id,
	}
	r := bufio.NewReader(bytes.NewReader(input))
	if _, err := r.ReadBytes('\n'); err != nil {
		return data, err
	}
	header, err := r.ReadBytes('\n')
	if err != nil {
		return data, err
	}
	if !p.header.Match(header) {
		return data, fmt.Errorf("Parse: header doesn't match: id=%s, header=%s", Id, header)
	}
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 1 {
					if record, err := p.parseRecord(line); err == nil {
						data.Records = append(data.Records, record)
					}
				}
				return data, nil
			}
			return data, err
		}
		if record, err := p.parseRecord(line); err == nil {
			data.Records = append(data.Records, record)
		}
	}
}

func (p *parser) parseRecord(input []byte) (device.Record, error) {
	matches := p.re.FindSubmatch(input)
	if matches == nil {
		return device.Record{}, fmt.Errorf("parseRecord: invalid record: id=%s, input=%s", Id, input)
	}
	return device.Record{
		Port:        string(matches[1]),
		Description: string(matches[2]),
		VLAN:        string(matches[3]),
	}, nil
}
