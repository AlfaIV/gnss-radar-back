package middlewarecustom

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	auth_domain "gnss-radar/gnss-api-gateway/internal/auth"
)

type UserIDMiddleware struct {
	authUsecase auth_domain.Usecase
	logger      *logrus.Logger
}

func NewUserIDMW(au auth_domain.Usecase, logger *logrus.Logger) *UserIDMiddleware {
	return &UserIDMiddleware{
		authUsecase: au,
		logger:      logger,
	}
}

func (mw *UserIDMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionId, err := c.Cookie(auth_domain.CookieName)
		if err != nil {
			mw.logger.Error("userID middleware: ", err)
			return c.String(http.StatusUnauthorized, "failed to get session id")
		}

		userID, err := mw.authUsecase.GetUserId(c.Request().Context(), sessionId.Value)
		if err != nil {
			mw.logger.Error("userID middleware: ", err)
			return c.String(http.StatusUnauthorized, "invalid session id")
		}

		ctx := context.WithValue(c.Request().Context(), UserIDKey, userID)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
