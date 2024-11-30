package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SP3FirstOnePlusLine struct {
	NumberOfSats int   `json:"number_of_sats"`
	Ids          SatId `json:"ids"`
}

type SatId struct {
	Symbol byte
	Ids    []string
}

func ParseSP3FirstOnePlusLine(line string) (SP3FirstOnePlusLine, error) {
	line = removeExtraSpaces(line)
	var thirdLine SP3FirstOnePlusLine

	if line[0:1] != "+" {
		return thirdLine, errors.New("invalid format of SP3 third line")
	}

	thirdLine.NumberOfSats = parseInteger(line[2:5])

	satIds := parseSatIds(line[6:57], 'C')
	thirdLine.Ids = satIds

	return thirdLine, nil
}

func parseInteger(input string) int {
	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Error parsing integer: %v\n", err)
		return 0
	}
	return value
}

func parseSatIds(input string, symbol byte) SatId {
	satId := SatId{
		Symbol: symbol,
		Ids:    make([]string, 0, len(input)),
	}
	Ids := strings.Split(input, string(symbol))
	for _, id := range Ids {
		if id != "" {
			satId.Ids = append(satId.Ids, id)
		}
	}
	return satId
}
