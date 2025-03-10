package measurements_calculate

import (
	"math"
	"time"

	"github.com/joshuaferrara/go-satellite"
)

type Coordinates struct {
	XY satellite.LatLong
	Z  float64
}
type TLEFile struct {
	tleLine1, tleLine2 string
}

func DeterminingCoordinates(coordinates Coordinates, tlefile TLEFile) (float64, float64) {

	sat := satellite.TLEToSat(tlefile.tleLine1, tlefile.tleLine2, "wgs72")

	now := time.Now()

	position, _ := satellite.Propagate(sat, now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())

	observerECI := satellite.LLAToECI(coordinates.XY, coordinates.Z, satellite.JDay(now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second()))

	dx := position.X - observerECI.X
	dy := position.Y - observerECI.Y
	dz := position.Z - observerECI.Z

	observerLat := coordinates.XY.Latitude
	observerLon := coordinates.XY.Longitude
	azimuth, elevation := calculateLookAngles(observerLat, observerLon, dx, dy, dz)

	return azimuth * 180 / math.Pi, elevation * 180 / math.Pi
}

func calculateLookAngles(observerLat, observerLon, dx, dy, dz float64) (float64, float64) {
	observerLatRad := observerLat * math.Pi / 180
	observerLonRad := observerLon * math.Pi / 180

	sinLat := math.Sin(observerLatRad)
	cosLat := math.Cos(observerLatRad)
	sinLon := math.Sin(observerLonRad)
	cosLon := math.Cos(observerLonRad)

	east := -sinLon*dx + cosLon*dy
	north := -sinLat*cosLon*dx - sinLat*sinLon*dy + cosLat*dz
	up := cosLat*cosLon*dx + cosLat*sinLon*dy + sinLat*dz

	azimuth := math.Atan2(east, north)
	if azimuth < 0 {
		azimuth += 2 * math.Pi
	}

	rangeSat := math.Sqrt(east*east + north*north + up*up)
	elevation := math.Asin(up / rangeSat)

	return azimuth, elevation
}

/*
func main() {

	coordinates := Coordinates{
		XY: satellite.LatLong{
			Latitude:  55.7558,
			Longitude: 37.7591,
		},
		Z: 120.0,
	}

	tlefile := TLEFile{
		tleLine1: "1 24876U 97035A   25067.52663765 -.00000009  00000+0  00000+0 0  9999",
		tleLine2: "2 24876  55.7641 116.7908 0087302  53.9556 306.9171  2.00562536202619",
	}

	azimuth, elevation := DeterminingCoordinates(coordinates, tlefile)

	fmt.Printf("Азимут: %.2f°\n", azimuth)
	fmt.Printf("Угол места: %.2f°\n", elevation)
}
*/
