package user_server

import (
	"bytes"
	"context"
	common_proto "gnss-radar/api/proto/common"
	proto "gnss-radar/api/proto/measurements"
	measurements_domain "gnss-radar/gnss-measurements/internal"
	"net/http"
	"path"

	google_proto "github.com/golang/protobuf/ptypes/empty"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MeasurementsServiceServer struct {
	logger *logrus.Logger
	repo   measurements_domain.Repository

	proto.UnimplementedMeasurementsServer
}

func NewMeasurementsServer(repo measurements_domain.Repository, logger *logrus.Logger) MeasurementsServiceServer {
	return MeasurementsServiceServer{repo: repo, logger: logger}
}

func (s *MeasurementsServiceServer) GetEphemeris(ctx context.Context, in *common_proto.PaginatedRequest) (*proto.Ephemeris, error) {

	fileMetas, total, err := s.repo.GetEphemeris(ctx, measurements_domain.PaginatedRequest{Page: in.Page, Size: in.Size})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "[MEASUREMENTS]: %v", err)
	}

	var ephemeris []*proto.EphemerisFileMeta
	for _, meta := range fileMetas {
		ephemeris = append(ephemeris, &proto.EphemerisFileMeta{
			Name:     meta.Name,
			Datetime: meta.Datetime,
		})
	}

	return &proto.Ephemeris{
		Ephemeris: ephemeris,
		Total:     total,
	}, nil
}

func (s *MeasurementsServiceServer) LoadEphemeris(ctx context.Context, in *proto.EphemerisToLoad) (*google_proto.Empty, error) {
	payload := in.GetPayload()
	contentType := http.DetectContentType(payload)

	file := measurements_domain.EphemerisToLoad{
		Name:        in.Name,
		MinioName:   uuid.New().String() + "." + path.Base(contentType),
		Payload:     bytes.NewReader(payload),
		PayloadSize: int64(len(payload)),
		ContentType: contentType,
	}

	err := s.repo.UploadEphemeris(ctx, file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to upload avatar")
	}

	return &google_proto.Empty{}, nil
}

func (s *MeasurementsServiceServer) GetSatellitesPosition(ctx context.Context, in *google_proto.Empty) (*proto.Satellites, error) {

	satellites, err := s.repo.GetSatellitesCoordinates(ctx)
	if err != nil {
		return nil, err
	}

	var satArray []*proto.Satellite
	for _, sat := range satellites.Satellites {
		satArray = append(satArray, &proto.Satellite{
			Name:     sat.Name,
			Group: sat.Group,
			Azimuth: sat.Azumuth,
			Elevation: sat.Elevation,
			Range: sat.Range,
		})
	}

	return &proto.Satellites{
		Satellites: satArray,
	}, nil
}