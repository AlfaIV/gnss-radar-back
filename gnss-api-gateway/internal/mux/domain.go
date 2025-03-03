package mux

import (
	auth_domain_gateway "gnss-radar/gnss-api-gateway/internal/auth"
	auth_handler "gnss-radar/gnss-api-gateway/internal/auth/delivery"
	"gnss-radar/gnss-api-gateway/internal/config"
	middlewarecustom "gnss-radar/gnss-api-gateway/internal/mux/middleware"
	user_domain_gateway "gnss-radar/gnss-api-gateway/internal/user"
	user_handler "gnss-radar/gnss-api-gateway/internal/user/delivery"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	Auth auth_handler.AuthHandler
	User user_handler.UserHandler
}

type ServiceUsecase struct {
	Auth auth_domain_gateway.Usecase
	User user_domain_gateway.Usecase
}

func Setup(config *config.Config, service ServiceUsecase, handlers Handlers, logger *logrus.Logger) *echo.Echo {
	mux := echo.New()

	userIDMiddleware := middlewarecustom.NewUserIDMW(service.Auth, logger)
	userPermissionsMiddleware := middlewarecustom.NewUserPermissionsMiddleware(service.User, logger)

	mux.Use(middleware.Recover())
	mux.Use(middleware.CORSWithConfig(config.CORS))
	mux.Use(middleware.RequestID())

	base := mux.Group("/api/v1")

	auth := base.Group("/auth")
	auth.POST("/login", handlers.Auth.Login) // api/v1/auth/login
	auth.POST("/signup", handlers.Auth.Signup)
	auth.DELETE("/logout", handlers.Auth.Logout)
	auth.GET("/me", handlers.Auth.Me)

	user := base.Group("/user")
	user.Use(
		userIDMiddleware.Process,
		userPermissionsMiddleware.Process,
	)
	user.GET("/getListUsers", handlers.User.GetListUsers)
	user.GET("/getSignUpRequestions", handlers.User.GetSignUpRequestions)
	user.PATCH("/resolveSignUp", handlers.User.ResolveUserSignUp)
	user.PATCH("/givePermissions", handlers.User.GivePermissions)

	return mux
}
