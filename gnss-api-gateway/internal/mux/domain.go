package mux

import (
	auth_domain "gnss-radar/gnss-api-gateway/internal/auth"
	"gnss-radar/gnss-api-gateway/internal/config"
	middlewarecustom "gnss-radar/gnss-api-gateway/internal/mux/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func Setup(config *config.Config, authUsecase auth_domain.Usecase, logger *logrus.Logger) *echo.Echo {
	mux := echo.New()

	userIDMiddleware := middlewarecustom.NewUserIDMW(authUsecase, logger)

	mux.Use(middleware.Recover())
	mux.Use(middleware.CORSWithConfig(config.CORS))
	mux.Use(middleware.RequestID())

	base := mux.Group("/api/v1")

	_ = base.Group("/auth")

	user := base.Group("/user")
	user.Use(userIDMiddleware.Process)


	return mux
}
