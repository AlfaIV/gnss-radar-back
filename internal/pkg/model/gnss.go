package model

import "time"

type Coords struct {
	X float64 `db:"x"`
	Y float64 `db:"y"`
	Z float64 `db:"z"`
}

type GnssCoords struct {
	ID            string  `db:"id"`
	SatelliteID   string  `db:"satellite_id"`
	SatelliteName string  `db:"satellite_name"`
	X             float64 `db:"x"`
	Y             float64 `db:"y"`
	Z             float64 `db:"z"`
}

type Device struct {
	ID          string  `db:"id"`
	Name        string  `db:"name"`
	Token       string  `db:"token"`
	Description *string `db:"description"`
	X           float64 `db:"x"`
	Y           float64 `db:"y"`
	Z           float64 `db:"z"`
}

type Task struct {
	ID           string       `db:"id"`
	SatelliteID  string       `db:"satellite_id"`
	SignalType   SignalType   `db:"signal_type"`
	GroupingType GroupingType `db:"grouping_type"`
	StartAt      time.Time    `db:"start_at"`
	EndAt        time.Time    `db:"end_at"`
}
