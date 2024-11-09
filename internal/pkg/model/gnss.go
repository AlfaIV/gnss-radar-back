package model

type Coords struct {
	X float64 `db:"x"`
	Y float64 `db:"y"`
	Z float64 `db:"z"`
}

type GnssCoords struct {
	ID            string `db:"id"`
	SatelliteID   string `db:"satellite_id"`
	SatelliteName string `db:"satellite_name"`
	Coords        Coords
}

type Device struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Token       string `db:"token"`
	Description string `db:"description"`
	Coords      Coords
}
