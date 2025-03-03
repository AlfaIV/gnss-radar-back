package middlewarecustom

import (
	user_domain_gateway "gnss-radar/gnss-api-gateway/internal/user"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserPermissionsMiddleware struct {
	userUsecase user_domain_gateway.Usecase
	logger      *logrus.Logger
}

func NewUserPermissionsMiddleware(uu user_domain_gateway.Usecase, logger *logrus.Logger) *UserPermissionsMiddleware {
	return &UserPermissionsMiddleware{
		userUsecase: uu,
		logger:      logger,
	}
}

func (mw *UserPermissionsMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, ok := c.Request().Context().Value(UserIDKey).(string)
		if !ok || userID == "" {
			mw.logger.Error("userPermissions middleware: userID not found in context")
			return c.String(http.StatusUnauthorized, "user identification failed")
		}

		fullPath := c.Path()
		pathParts := strings.Split(fullPath, "/")
		
		var method string
		for i := len(pathParts) - 1; i >= 0; i-- {
			if pathParts[i] != "" {
				method = pathParts[i]
				break
			}
		}
		
		if method == "" {
			mw.logger.Error("userPermissions middleware: empty method name")
			return c.String(http.StatusBadRequest, "invalid endpoint")
		}

		status, err := mw.userUsecase.ValidatePermissions(c.Request().Context(), userID, method)
		if err != nil || !status {
			mw.logger.WithFields(logrus.Fields{
				"user_id": userID,
				"method":  method,
			}).Error("permission check failed: ", err)
			
			return c.String(http.StatusForbidden, "access denied")
		}

		return next(c)
	}
}