package utils

import (
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

func SerializerGnssCoords(list []*model.GnssCoords) []*model.Gnss {
	return lo.Map(list, func(item *model.GnssCoords, _ int) *model.Gnss {
		return &model.Gnss{
			ID:            item.ID,
			SatelliteID:   item.SatelliteID,
			SatelliteName: item.SatelliteName,
			Coordinates: &model.CoordsResults{
				X: strconv.FormatFloat(item.Coords.X, 'f', -1, 64),
				Y: strconv.FormatFloat(item.Coords.Y, 'f', -1, 64),
				Z: strconv.FormatFloat(item.Coords.Z, 'f', -1, 64),
			},
		}
	})
}
