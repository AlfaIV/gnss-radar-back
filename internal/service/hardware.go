package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/pkg/parser"
	"github.com/Gokert/gnss-radar/internal/store"
)

type IHardware interface {
	AddSpectrum(ctx context.Context, spectrumReq model.SpectrumRequest) error
	AddPower(ctx context.Context, powerReq model.PowerRequest) error
	AddPairMeasurement(ctx context.Context, pairMeasurementReq model.PairMeasurementRequest) error
	CompareDeviceToken(ctx context.Context, deviceToken string) error
	UploadSP3(ctx context.Context, pathWithFiles string) error
}

type Hardware struct {
	store store.IGnssStore
}

func NewHardwareService(store store.IGnssStore) *Hardware {
	return &Hardware{store: store}
}

func (h *Hardware) AddSpectrum(ctx context.Context, spectrumReq model.SpectrumRequest) error {
	if err := h.store.AddSpectrum(ctx, spectrumReq); err != nil {
		return fmt.Errorf("h.store.AddSpectrum: %w", err)
	}
	return nil
}

func (h *Hardware) AddPower(ctx context.Context, powerReq model.PowerRequest) error {
	if err := h.store.AddPower(ctx, powerReq); err != nil {
		return fmt.Errorf("h.store.AddPower: %w", err)
	}
	return nil
}

func (h *Hardware) CompareDeviceToken(ctx context.Context, deviceToken string) error {
	if err := h.store.CompareDeviceToken(ctx, deviceToken); err != nil {
		return fmt.Errorf("h.store.CompareDeviceToken: %w", err)
	}
	return nil
}

func (h *Hardware) AddPairMeasurement(ctx context.Context, pairMeasurementReq model.PairMeasurementRequest) error {
	if err := h.store.AddPairMeasurement(ctx, pairMeasurementReq); err != nil {
		return fmt.Errorf("h.store.AddPairMeasurement: %w", err)
	}
	return nil
}

func (h *Hardware) UploadSP3(ctx context.Context, pathWithFiles string) error {
	files, err := os.ReadDir(pathWithFiles)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, file := range files {
		filePath := filepath.Join(pathWithFiles, file.Name())

		parsedFile, err := parser.ParseSP3File(filePath)
		if err != nil {
			return fmt.Errorf("failed to parse SP3 file %s: %w", filePath, err)
		}

		for _, satellite := range parsedFile.SatelliteLines[:len(parsedFile.SatelliteLines)-1] {
			x, y, z, err := parseCoordinates(satellite.CoordinateSystem)
			if err != nil {
				return fmt.Errorf("failed to parse coordinates for satellite %s: %w", satellite.SatelliteId, err)
			}

			err = h.store.SaveParsedSP3(ctx, satellite.SatelliteId, x, y, z, parsedFile.TimeLines[satellite.TimeLineId-1])
			if err != nil {
				return fmt.Errorf("failed to save parsed SP3 for satellite %s: %w", satellite.SatelliteId, err)
			}
		}
	}

	return nil
}

func parseCoordinates(coordStr string) (float64, float64, float64, error) {
	parts := strings.Fields(coordStr)
	fmt.Println(parts, coordStr)

	if len(parts) < 3 {
		return 0, 0, 0, fmt.Errorf("invalid input: not enough parts")
	}

	x, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse x coordinate: %w", err)
	}

	y, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse y coordinate: %w", err)
	}

	z, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse z coordinate: %w", err)
	}

	return x, y, z, nil
}
