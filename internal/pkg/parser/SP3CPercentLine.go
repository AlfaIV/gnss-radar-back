package parser

type SP3CPercentLine struct {
	FileType   string `json:"file_type"`
	SystemType string `json:"system_type"`
}

func ParseSP3CPercentLine(line string) (SP3CPercentLine, error) {
	line = removeExtraSpaces(line)
	parts := fileAndSystemType(line)
	if len(parts) == 0 {
		return SP3CPercentLine{}, nil
	}

	return SP3CPercentLine{
		FileType:   string(parts[0]),
		SystemType: string(parts[2:]),
	}, nil
}

func fileAndSystemType(line string) string {
	fileAndSystem := ""

	for i := 0; i < len(line); i++ {
		if line[i] != '%' && line[i] != 'c' {
			fileAndSystem += string(line[i])
		}
	}
	fileAndSystem = removeExtraSpaces(fileAndSystem)
	return fileAndSystem
}
