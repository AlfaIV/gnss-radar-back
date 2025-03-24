package measurements_client

import (
	"context"
	common_proto "gnss-radar/api/proto/common"
	proto "gnss-radar/api/proto/measurements"
	measurements_domain_gateway "gnss-radar/gnss-api-gateway/internal/measurements"
	"io"

	google_proto "github.com/golang/protobuf/ptypes/empty"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type MeasurementsClient struct {
	client proto.MeasurementsClient
	logger *logrus.Logger
}

func NewMeasurementsClient(client proto.MeasurementsClient, logger *logrus.Logger) MeasurementsClient {
	return MeasurementsClient{client: client, logger: logger}
}

func (mc *MeasurementsClient) GetEphemeris(ctx context.Context, page uint64, size uint64) ([]measurements_domain_gateway.EphemerisFileMeta, uint64, error) {
	ephemeris, err := mc.client.GetEphemeris(ctx, &common_proto.PaginatedRequest{Page: page, Size: size})
	if err != nil {
		return []measurements_domain_gateway.EphemerisFileMeta{}, 0, errors.Wrapf(err, "[GW USER] %v", err)
	}

	var ephemerisArray []measurements_domain_gateway.EphemerisFileMeta

	for _, eph := range ephemeris.Ephemeris {
		ephemerisArray = append(ephemerisArray, measurements_domain_gateway.EphemerisFileMeta{
			Name:     eph.Name,
			Datetime: eph.Datetime,
		})
	}

	return ephemerisArray, ephemeris.GetTotal(), nil
}

func (mc *MeasurementsClient) LoadEphemeris(ctx context.Context, file io.Reader, name string) error {
	payload, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	_, err = mc.client.LoadEphemeris(ctx, &proto.EphemerisToLoad{
		Name:    name,
		Payload: payload,
	})
	if err != nil {
		return errors.Wrapf(err, "[GW MEASUREMENTS] %v", err)
	}

	return nil
}

func (mc *MeasurementsClient) GetSatellitesPosition(ctx context.Context) (measurements_domain_gateway.Satellites, error) {
	s, err := mc.client.GetSatellitesPosition(ctx, &google_proto.Empty{})
	if err != nil {
		return measurements_domain_gateway.Satellites{}, errors.Wrapf(err, "[GW MEASUREMENTS] %v", err)
	}

	var satellites []measurements_domain_gateway.SatelliteData

	for _, sat := range s.GetSatellites() {
		satellites = append(satellites, measurements_domain_gateway.SatelliteData{
			Group: sat.GetGroup(),
			Name: sat.GetName(),
			Azumuth: sat.GetAzimuth(),
			Elevation: sat.GetElevation(),
			Range: sat.GetRange(),
		})
	}

	return measurements_domain_gateway.Satellites{Satellites: satellites}, nil
}