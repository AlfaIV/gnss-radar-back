package config

import (
	"gnss-radar/gnss-api-gateway/internal/config/cors"

	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	CORS middleware.CORSConfig
}

func NewConfig() (*Config, error) {
	corsConfig, err := cors.New()

	if err != nil {
		return nil, err
	}

	return &Config{
		CORS: corsConfig,
	}, nil
}
