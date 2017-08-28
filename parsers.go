// Package msix_parsers importing concrete implementations of github.com/doctor-fate/mskix.Parser
// and create aliases for Id's
// End users must import this package instead of concrete implementations
package mskix_parsers

import (
	"github.com/doctor-fate/mskix-parsers/cisco"
	"github.com/doctor-fate/mskix-parsers/extreme"
	"github.com/doctor-fate/mskix-parsers/force10"
)

const (
	Force10 = force10.Id
	Cisco   = cisco.Id
	Extreme = extreme.Id
)
