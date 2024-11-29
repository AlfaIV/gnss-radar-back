package parser

type SP3SatelliteLine struct {
	TimeLineId       int    `json:"time_line_id"`
	SatelliteId      string `json:"satellite_id"`
	CoordinateSystem string `json:"coordinate_system"`
}

func ParseSP3SatelliteLine(line string, timeLineId int) (SP3SatelliteLine, error) {
	line = removeExtraSpaces(line)

	satelliteId := ""
	for i := 0; i < len(line); i++ {
		if line[i] != ' ' && line[i] != '\n' {
			satelliteId += string(line[i])
		} else {
			break
		}
	}

	return SP3SatelliteLine{TimeLineId: timeLineId, SatelliteId: satelliteId, CoordinateSystem: line[len(satelliteId):]}, nil
}
