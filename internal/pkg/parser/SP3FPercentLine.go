package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SP3FPercentLine struct {
	FLine []float64 `json:"f_line"`
}

func ParseSP3FPercentLine(line string) (SP3FPercentLine, error) {
	line = removeExtraSpaces(line)
	var fPercentLine SP3FPercentLine

	if line[0:2] != "%f" {
		return fPercentLine, errors.New("invalid format of SP3 F percent line")
	}

	values := strings.Fields(line[2:])
	if len(values) != 4 {
		return fPercentLine, errors.New("invalid number of values in SP3 F percent line")
	}

	for _, value := range values {
		fValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fPercentLine, fmt.Errorf("invalid float value: %s", value)
		}
		fPercentLine.FLine = append(fPercentLine.FLine, fValue)
	}

	return fPercentLine, nil
}
