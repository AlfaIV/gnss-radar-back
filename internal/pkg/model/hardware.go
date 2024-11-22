package model

import "time"

type Description struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Group     string    `json:"Group"`
	Signal    string    `json:"Signal"`
	Target    string    `json:"Target"`
}

type SpectrumData struct {
	Spectrum  []float32 `json:"spectrum"`
	StartFreq float32   `json:"StartFreq"`
	FreqStep  float32   `json:"FreqStep"`
	StartTime time.Time `json:"startTime"`
}

type PowerData struct {
	Power     []float32 `json:"power"`
	StartTime time.Time `json:"startTime"`
	TimeStep  time.Time `json:"timeStep"`
}

type SpectrumRequest struct {
	Token       string       `json:"token"`
	Description Description  `json:"description"`
	Data        SpectrumData `json:"data"`
}

type PowerRequest struct {
	Token       string      `json:"token"`
	Description Description `json:"description"`
	Data        PowerData   `json:"data"`
}
