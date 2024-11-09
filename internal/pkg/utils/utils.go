package utils

import (
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/store"
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

func SerializerGnssCoords(list []*store.ListResult) []*model.Gnss {
	return lo.Map(list, func(item *store.ListResult, _ int) *model.Gnss {
		return &model.Gnss{
			ID:            item.ID,
			SatelliteID:   item.SatelliteID,
			SatelliteName: item.SatelliteName,
			Coordinates: &model.CoordsResults{
				X: strconv.FormatFloat(item.X, 'f', -1, 64),
				Y: strconv.FormatFloat(item.Y, 'f', -1, 64),
				Z: strconv.FormatFloat(item.Z, 'f', -1, 64),
			},
		}
	})
}
