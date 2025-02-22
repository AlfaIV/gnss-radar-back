package cors

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v3"
)

func New() (middleware.CORSConfig, error) {
	var config middleware.CORSConfig

	data, err := os.ReadFile("config/cors.yml")
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	config.Skipper = middleware.DefaultSkipper

	return config, nil
}
