package parser

import (
	"errors"
)

type SP3OnePlusLine struct {
	OnePlusLine SatId `json:"one_plus_line"`
}

func ParseSP3OnePlusLine(line string) (SP3OnePlusLine, error) {
	line = removeExtraSpaces(line)
	var thirdLine SP3OnePlusLine

	if line[0:1] != "+" {
		return thirdLine, errors.New("invalid format of SP3 OnePlusLine")
	}

	satIds := parseSatIds(line[2:53], 'C')
	thirdLine.OnePlusLine = satIds

	return thirdLine, nil
}
