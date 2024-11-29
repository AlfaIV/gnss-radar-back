package parser

import (
	"errors"
	"strconv"
	"strings"
)

type SP3TimeLine struct {
	Id              int     `json:"id"`
	YearStart       int     `json:"year_start"`
	MonthStart      int     `json:"month_start"`
	DayOfMonthStart int     `json:"day_of_month_start"`
	HourStart       int     `json:"hour_start"`
	MinuteStart     int     `json:"minute_start"`
	SecondStart     float64 `json:"second_start"`
}

func ParseSP3TimeLine(line string, id int) (SP3TimeLine, error) {
	line = removeExtraSpaces(line)
	var timeLine SP3TimeLine

	if line[0:1] != "*" {
		return timeLine, errors.New("invalid format of SP3 time line")
	}

	parts := strings.Fields(line[1:])
	if len(parts) < 6 {
		return timeLine, errors.New("not enough data in SP3 time line")
	}

	var err error
	timeLine.Id = id
	timeLine.YearStart, err = strconv.Atoi(parts[0])
	if err != nil {
		return timeLine, err
	}
	timeLine.MonthStart, err = strconv.Atoi(parts[1])
	if err != nil {
		return timeLine, err
	}
	timeLine.DayOfMonthStart, err = strconv.Atoi(parts[2])
	if err != nil {
		return timeLine, err
	}
	timeLine.HourStart, err = strconv.Atoi(parts[3])
	if err != nil {
		return timeLine, err
	}
	timeLine.MinuteStart, err = strconv.Atoi(parts[4])
	if err != nil {
		return timeLine, err
	}
	timeLine.SecondStart, err = strconv.ParseFloat(parts[5], 64)
	if err != nil {
		return timeLine, err
	}

	return timeLine, nil
}
