package model

import "time"

type Coords struct {
	X float64 `db:"x"`
	Y float64 `db:"y"`
	Z float64 `db:"z"`
}

type GnssCoords struct {
	ID          string    `db:"id"`
	SatelliteID string    `db:"satellite_id"`
	X           float64   `db:"x"`
	Y           float64   `db:"y"`
	Z           float64   `db:"z"`
	CreatedAt   time.Time `db:"created_at"`
}

type Device struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Token       string    `db:"token"`
	Description *string   `db:"description"`
	X           float64   `db:"x"`
	Y           float64   `db:"y"`
	Z           float64   `db:"z"`
	CreatedAt   time.Time `db:"created_at"`
}

type Task struct {
	ID           string       `db:"id"`
	Title        string       `db:"title"`
	Description  *string      `db:"description"`
	SatelliteID  string       `db:"satellite_id"`
	DeviceID     string       `db:"device_id"`
	SignalType   SignalType   `db:"signal_type"`
	GroupingType GroupingType `db:"grouping_type"`
	StartAt      time.Time    `db:"start_at"`
	EndAt        time.Time    `db:"end_at"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
}

type SatelliteInfo struct {
	ID                        string    `db:"id"`
	ExternalSatelliteId       string    `db:"external_satellite_id"`
	SatelliteName             string    `db:"satellite_name"`
	CoordinateMeasurementTime time.Time `db:"coordinate_measurement_time"`
	CreatedAt                 time.Time `db:"created_at"`
}
