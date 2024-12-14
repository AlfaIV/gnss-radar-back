package utils

import (
	"crypto/sha512"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/samber/lo"
	"math/rand"
	"strconv"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(seed int) string {
	symbols := make([]rune, seed)
	for i := range symbols {
		symbols[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(symbols)
}

func HashPassword(password string) []byte {
	hashPassword := sha512.Sum512([]byte(password))
	passwordByteSlice := hashPassword[:]
	return passwordByteSlice
}

func SerializerGnssCoords(list []*model.GnssCoords) []*model.Gnss {
	return lo.Map(list, func(item *model.GnssCoords, _ int) *model.Gnss {
		return &model.Gnss{
			ID:          item.ID,
			SatelliteID: item.SatelliteID,
			Coordinates: &model.CoordsResults{
				X: strconv.FormatFloat(item.X, 'f', -1, 64),
				Y: strconv.FormatFloat(item.Y, 'f', -1, 64),
				Z: strconv.FormatFloat(item.Z, 'f', -1, 64),
			},
			CreatedAt:      item.CreatedAt,
			Azimuth:        rand.Intn(361),
			ElevationAngle: rand.Intn(361),
			Distance:       rand.Intn(100000),
		}
	})
}
