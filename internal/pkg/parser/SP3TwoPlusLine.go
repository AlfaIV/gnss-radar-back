package parser

import (
	"errors"
	"strconv"
)

type SP3TwoPlusLine struct {
	TwoPlusLine []int `json:"two_plus_line"`
}

func ParseSP3TwoPlusLine(line string) (SP3TwoPlusLine, error) {
	line = removeExtraSpaces(line)
	var twoPlusLine SP3TwoPlusLine
	if line[0:2] != "++" {
		return twoPlusLine, errors.New("invalid format of SP3 OnePlusLine")
	}

	twoLineInt, err := stringToIntegers(line[2:])
	if err != nil {
		return twoPlusLine, err
	}

	twoPlusLine.TwoPlusLine = twoLineInt
	return twoPlusLine, nil
}

func stringToIntegers(input string) ([]int, error) {
	numbers := []int{}
	buffer := ""

	for _, char := range input {
		if char == ' ' || char == '\n' {
			if buffer != "" {
				number, err := strconv.Atoi(buffer)
				if err != nil {
					return nil, err
				}
				numbers = append(numbers, number)
				buffer = ""
			}
		} else {
			buffer += string(char)
		}
	}

	if buffer != "" {
		number, err := strconv.Atoi(buffer)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}
