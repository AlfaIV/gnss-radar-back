package auth_handler

import (
	auth_domain_gateway "gnss-radar/gnss-api-gateway/internal/auth"
	user_domain_gateway "gnss-radar/gnss-api-gateway/internal/user"
	"net/http"
	"net/mail"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	userUsecase    user_domain_gateway.Usecase
	authUsecase    auth_domain_gateway.Usecase
	logger         *logrus.Logger
}

func NewHandler(
	user user_domain_gateway.Usecase,
	auth auth_domain_gateway.Usecase,
	logger *logrus.Logger,
) AuthHandler {
	return AuthHandler{
		userUsecase:    user,
		authUsecase:    auth,
		logger:         logger,
	}
}

func makeCookieWithSessionId(sessionId string) http.Cookie {
	var cookie http.Cookie
	cookie.Name = auth_domain_gateway.CookieName
	cookie.Value = sessionId
	cookie.Expires = time.Now().Add(auth_domain_gateway.CookieTTL)
	cookie.HttpOnly = true
	cookie.Path = "/api/v1"

	return cookie
}

func (h *AuthHandler) Login(c echo.Context) error {
	loginRequest := user_domain_gateway.LoginRequest{}
	if err := c.Bind(&loginRequest); err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusBadRequest, "failed to parse request data")
	}

	user, err := h.userUsecase.Login(c.Request().Context(), loginRequest.Login, loginRequest.Password)
	if err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusUnauthorized, "invalid data")
	}

	if user.Status == "PENDING" {
		return c.String(http.StatusForbidden, "forbidden")
	}

	sessionId, err := h.authUsecase.CreateSession(c.Request().Context(), user.Id)
	if err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusInternalServerError, "failed to create session")
	}

	cookie := makeCookieWithSessionId(sessionId)
	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) Signup(c echo.Context) error {

	user := user_domain_gateway.SignUpRequest{}
	if err := c.Bind(&user); err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusBadRequest, "failed to parse request data")
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusUnauthorized, "failed to validate email")
	}

	ctx := c.Request().Context()

	isOk, err := h.userUsecase.SignUp(ctx, user)
	if err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusInternalServerError, "failed to create account")
	}

	if !isOk {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusUnauthorized, "failed to create account")
	}

	return c.NoContent(http.StatusOK)
}

func (h *AuthHandler) Logout(c echo.Context) error {
	sessionId, err := c.Cookie(auth_domain_gateway.CookieName)
	if err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusUnauthorized, "failed to get session id")
	}

	if err := h.authUsecase.DeleteSession(c.Request().Context(), sessionId.Value); err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusBadRequest, "failed to delete session id")
	}

	return c.NoContent(http.StatusOK)
}

func (h *AuthHandler) Me(c echo.Context) error {
	sessionId, err := c.Cookie(auth_domain_gateway.CookieName)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	userId, err := h.authUsecase.GetUserId(c.Request().Context(), sessionId.Value)
	if err != nil {

		return c.String(http.StatusUnauthorized, "Failed to authenticate")
	}

	userInfo, err := h.userUsecase.GetUserInfoById(c.Request().Context(), userId)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.JSON(http.StatusInternalServerError, "Failed to get user info")
	}

	return c.JSON(http.StatusOK, userInfo)
}

