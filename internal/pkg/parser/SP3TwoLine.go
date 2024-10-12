package parser

import (
	"errors"
	"fmt"
	"strconv"
)

type SP3SecondLine struct {
	GPSWeek        int     `json:"gps_week"`
	SecondsOfWeek  float64 `json:"seconds_of_week"`
	EpochInterval  float64 `json:"epoch_interval"`
	ModJulDayStart int     `json:"mod_jul_day_start"`
	FractionalDay  float64 `json:"fractional_day"`
}

func ParseSP3SecondLine(line string) (SP3SecondLine, error) {
	line = removeExtraSpaces(line)
	var secondLine SP3SecondLine

	if line[0:2] != "##" {
		return secondLine, errors.New("invalid format of SP3 second line")
	}

	secondLine.GPSWeek = parserInteger(line[3:7])
	secondLine.SecondsOfWeek = parserFloat(line[8:18])
	secondLine.EpochInterval = parserFloat(line[19:31])
	secondLine.ModJulDayStart = parserInteger(line[32:37])
	secondLine.FractionalDay = parserFloat(line[38:])

	return secondLine, nil
}

func parserInteger(input string) int {

	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Error parsing integer: %v\n", err)
		return 0
	}
	return value
}

func parserFloat(input string) float64 {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("Error parsing float: %v\n", err)
		return 0.0
	}
	return value
}
