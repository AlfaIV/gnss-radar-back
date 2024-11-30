package parser

import (
	"strconv"
	"strings"
)

type SP3FirstLine struct {
	VersionSymbol   string  `json:"version_symbol"`
	PosOrVelFlag    string  `json:"pos_or_vel_flag"`
	YearStart       int     `json:"year_start"`
	MonthStart      int     `json:"month_start"`
	DayOfMonthStart int     `json:"day_of_month_start"`
	HourStart       int     `json:"hour_start"`
	MinuteStart     int     `json:"minute_start"`
	SecondStart     float64 `json:"second_start"`
	NumberOfEpochs  int     `json:"number_of_epochs"`
	DataUsed        string  `json:"data_used"`
	CoordinateSys   string  `json:"coordinate_sys"`
	OrbitType       string  `json:"orbit_type"`
	Agency          string  `json:"agency"`
}

func ParseSP3FirstLine(line string) (SP3FirstLine, error) {
	line = removeExtraSpaces(line)
	var firstLine SP3FirstLine

	firstLine.VersionSymbol = line[0:2]
	firstLine.PosOrVelFlag = line[2:3]

	var err error
	firstLine.YearStart, err = parseInt(line[3:7])
	if err != nil {
		return firstLine, err
	}

	firstLine.MonthStart, err = parseInt(line[8:10])
	if err != nil {
		return firstLine, err
	}

	firstLine.DayOfMonthStart, err = parseNotSingleNumber(line[10:12])
	if err != nil {
		return firstLine, err
	}

	firstLine.HourStart, err = parseNotSingleNumber(line[14:16])
	if err != nil {
		return firstLine, err
	}

	firstLine.MinuteStart, err = parseNotSingleNumber(line[16:18])
	if err != nil {
		return firstLine, err
	}

	firstLine.SecondStart, err = parseFloat(line[20:27])
	if err != nil {
		return firstLine, err
	}

	numberOfEpochsStr := line[28:30]
	if numberOfEpochsStr != "" {
		firstLine.NumberOfEpochs, err = parseInt(numberOfEpochsStr)
		if err != nil {
			return firstLine, err
		}
	}

	firstLine.DataUsed = line[31:36]
	firstLine.CoordinateSys = line[37:42]
	firstLine.OrbitType = line[43:46]
	firstLine.Agency = line[47:50]

	return firstLine, nil
}

func removeExtraSpaces(input string) string {
	return strings.Join(strings.Fields(input), " ")
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(removeExtraSpaces(s))
}

func parseFloat(s string) (float64, error) {
	if len(removeExtraSpaces(s[0:2])) == 1 {
		return strconv.ParseFloat(s[1:], 64)
	}
	return strconv.ParseFloat(removeExtraSpaces(s), 64)
}

func parseNotSingleNumber(s string) (int, error) {
	if len(removeExtraSpaces(s)) == 2 {
		return strconv.Atoi(removeExtraSpaces(s))
	}
	if s[0] != ' ' {
		return strconv.Atoi(removeExtraSpaces(s[0:1]))
	}
	return strconv.Atoi(removeExtraSpaces(s[1:2]))
}
