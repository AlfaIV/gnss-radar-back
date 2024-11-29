package parser

import (
	"bufio"
	"fmt"
	"os"
)

const SkipThreeLines = 3

type SP3File struct {
	FirstLine        SP3FirstLine        `json:"first_line"`
	SecondLine       SP3SecondLine       `json:"second_line"`
	FirstOnePlusLine SP3FirstOnePlusLine `json:"first_one_plus_line"`
	OnePlusLines     []SP3OnePlusLine    `json:"one_plus_lines"`
	TwoPlusLines     []SP3TwoPlusLine    `json:"two_plus_lines"`
	CPercent         SP3CPercentLine     `json:"c_percent"`
	FPercentLines    []SP3FPercentLine   `json:"f_percent_lines"`
	IPercentLines    []SP3IPercentLine   `json:"i_percent_lines"`
	TimeLines        []SP3TimeLine       `json:"time_lines"`
	SatelliteLines   []SP3SatelliteLine  `json:"satellite_lines"`
}

func ParseSP3File(path string) (SP3File, error) {
	file, err := os.Open(path)
	if err != nil {
		return SP3File{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sp3File SP3File

	stringNumber := 0
	timelineid := 0
	for scanner.Scan() {
		line := scanner.Text()
		switch stringNumber {
		case 0:
			sp3File.FirstLine, err = ParseSP3FirstLine(line)
			if err != nil {
				return SP3File{}, fmt.Errorf("failed to parse first line: %w", err)
			}
		case 1:
			sp3File.SecondLine, err = ParseSP3SecondLine(line)
			if err != nil {
				return SP3File{}, fmt.Errorf("failed to parse second line: %w", err)
			}
		case 2:
			sp3File.FirstOnePlusLine, err = ParseSP3FirstOnePlusLine(line)
			if err != nil {
				return SP3File{}, fmt.Errorf("failed to parse first one plus line: %w", err)
			}
		}
		if stringNumber < SkipThreeLines {
			stringNumber++
		} else {
			switch line[0:2] {
			case "+ ":
				parsedLine, err := ParseSP3OnePlusLine(line)
				if err != nil {
					return SP3File{}, fmt.Errorf("failed to parse one plus line: %w", err)
				}
				sp3File.OnePlusLines = append(sp3File.OnePlusLines, parsedLine)
			case "++":
				parsedLine, err := ParseSP3TwoPlusLine(line)
				if err != nil {
					return SP3File{}, fmt.Errorf("failed to parse two plus line: %w", err)
				}
				sp3File.TwoPlusLines = append(sp3File.TwoPlusLines, parsedLine)
			case "%c":
				parsedLine, err := ParseSP3CPercentLine(line)
				if err != nil {
					return SP3File{}, fmt.Errorf("failed to parse C percent line: %w", err)
				}
				sp3File.CPercent = parsedLine
			case "%f":
				parsedLine, err := ParseSP3FPercentLine(line)
				if err != nil {
					return SP3File{}, fmt.Errorf("failed to parse F percent line: %w", err)
				}
				sp3File.FPercentLines = append(sp3File.FPercentLines, parsedLine)
			case "%i":
				parsedLine, err := ParseSP3IPercentLine(line)
				if err != nil {
					return SP3File{}, fmt.Errorf("failed to parse I percent line: %w", err)
				}
				sp3File.IPercentLines = append(sp3File.IPercentLines, parsedLine)
			case "/*":
			case "* ":
				timelineid++
				parsedLine, err := ParseSP3TimeLine(line, timelineid)
				if err != nil {
					return SP3File{}, fmt.Errorf("failed to parse time line: %w", err)
				}
				sp3File.TimeLines = append(sp3File.TimeLines, parsedLine)
			default:
				parsedLine, err := ParseSP3SatelliteLine(line, timelineid)
				if err != nil {
					return SP3File{}, fmt.Errorf("failed to parse satellite line: %w", err)
				}
				sp3File.SatelliteLines = append(sp3File.SatelliteLines, parsedLine)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return SP3File{}, fmt.Errorf("error reading file: %w", err)
	}
	return sp3File, nil
}
