package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SP3IPercentLine struct {
	IPercentLine []int `json:"i_percent_line"`
}

func ParseSP3IPercentLine(line string) (SP3IPercentLine, error) {
	line = removeExtraSpaces(line)
	var iPercentLine SP3IPercentLine

	if line[0:2] != "%i" {
		return iPercentLine, errors.New("invalid format of SP3 I percent line")
	}

	values := strings.Fields(line[2:])
	if len(values) != 9 {
		return iPercentLine, errors.New("invalid number of values in SP3 I percent line")
	}

	for _, value := range values {
		iValue, err := strconv.Atoi(value)
		if err != nil {
			return iPercentLine, fmt.Errorf("invalid integer value: %s", value)
		}
		iPercentLine.IPercentLine = append(iPercentLine.IPercentLine, iValue)
	}

	return iPercentLine, nil
}
