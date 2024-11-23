package service

import (
	"context"
	"fmt"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/store"
)

type IHardware interface {
	AddSpectrum(ctx context.Context, spectrumReq model.SpectrumRequest) error
	AddPower(ctx context.Context, powerReq model.PowerRequest) error
	CompareDeviceToken(ctx context.Context, deviceToken string) error
}

type Hardware struct {
	store store.IGnssStore
}

func NewHardwareService(store store.IGnssStore) *Hardware {
	return &Hardware{store: store}
}

func (h *Hardware) AddSpectrum(ctx context.Context, spectrumReq model.SpectrumRequest) error {
	err := h.store.AddSpectrum(ctx, spectrumReq)
	fmt.Println(spectrumReq, err)
	if err != nil {
		return fmt.Errorf("h.store.AddSpectrum: %w", err)
	}
	return nil
}

func (h *Hardware) AddPower(ctx context.Context, powerReq model.PowerRequest) error {
	err := h.store.AddPower(ctx, powerReq)
	fmt.Println(powerReq, err)
	if err != nil {
		return fmt.Errorf("h.store.AddPower: %w", err)
	}
	return nil
}

func (h *Hardware) CompareDeviceToken(ctx context.Context, deviceToken string) error {
	err := h.store.CompareDeviceToken(ctx, deviceToken)
	if err != nil {
		return fmt.Errorf("h.store.CompareDeviceToken: %w", err)
	}
	return nil
}
