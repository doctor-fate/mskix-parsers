package utility

import (
	"bufio"
	"io"

	"github.com/doctor-fate/mskix/device"
)

// ReadInput reads from r until EOF and invoke f for each line
func ReadInput(r *bufio.Reader, f func([]byte) (device.Record, error)) ([]device.Record, error) {
	var records []device.Record
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 1 {
					if record, err := f(line); err == nil {
						records = append(records, record)
					}
				}
				return records, nil
			}
			return records, err
		}
		if record, err := f(line); err == nil {
			records = append(records, record)
		}
	}
}
