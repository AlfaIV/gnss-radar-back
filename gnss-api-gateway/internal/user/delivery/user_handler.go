package user_handler

import (
	auth_domain_gateway "gnss-radar/gnss-api-gateway/internal/auth"
	user_domain_gateway "gnss-radar/gnss-api-gateway/internal/user"
	"gnss-radar/gnss-api-gateway/pkg/mwutils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userUsecase    user_domain_gateway.Usecase
	logger         *logrus.Logger
}

func NewHandler(
	user user_domain_gateway.Usecase,
	logger *logrus.Logger,
) UserHandler {
	return UserHandler{
		userUsecase:    user,
		logger:         logger,
	}
}

func (h *UserHandler) GetListUsers(c echo.Context) error {
	// Проверка куки

	ctx := c.Request().Context()

	_, err := mwutils.GetUserID(ctx)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusUnauthorized, "No session id provided")
	}

	//Проверка параметров
	pageParam := c.QueryParam("page")
	if pageParam == "" {

		return c.String(http.StatusBadRequest, "Incorrect page param")
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusBadRequest, "Incorrect page param")
	}

	sizeParam := c.QueryParam("size")
	if pageParam == "" {

		return c.String(http.StatusBadRequest, "Incorrect size param")
	}

	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusBadRequest, "Incorrect size param")
	}

	users, err := h.userUsecase.GetListUsers(c.Request().Context(), uint64(page), uint64(size))
	if err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusUnauthorized, "invalid data")
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetSignUpRequestions(c echo.Context) error {

	//Проверка куки

	ctx := c.Request().Context()

	_, err := mwutils.GetUserID(ctx)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusUnauthorized, "No session id provided")
	}

	// Проверка параметров
	pageParam := c.QueryParam("page")
	if pageParam == "" {

		return c.String(http.StatusBadRequest, "Incorrect page param")
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusBadRequest, "Incorrect page param")
	}

	sizeParam := c.QueryParam("size")
	if pageParam == "" {

		return c.String(http.StatusBadRequest, "Incorrect size param")
	}

	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusBadRequest, "Incorrect size param")
	}

	users, err := h.userUsecase.GetSignUpRequestions(c.Request().Context(), uint64(page), uint64(size))
	if err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusUnauthorized, "invalid data")
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) ResolveUserSignUp(c echo.Context) error {
	_, err := c.Cookie(auth_domain_gateway.CookieName)
	if err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusUnauthorized, "failed to get session id")
	}

	resolutionRequest := user_domain_gateway.SignUpResolutionRequest{}
	if err := c.Bind(&resolutionRequest); err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusBadRequest, "failed to parse request data")
	}

	if err := h.userUsecase.ResolveUserSignUp(c.Request().Context(), resolutionRequest.UserLogin, resolutionRequest.Resolution); err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusInternalServerError, "failed to make resolution")
	}

	return c.NoContent(http.StatusOK)
}

func (h *UserHandler) GivePermissions(c echo.Context) error {
	_, err := c.Cookie(auth_domain_gateway.CookieName)
	if err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusUnauthorized, "failed to get session id")
	}

	permissionChangeRequest := user_domain_gateway.PermissionChangeRequest{}
	if err := c.Bind(&permissionChangeRequest); err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusBadRequest, "failed to parse request data")
	}

	if err := h.userUsecase.ChangeUserPermissions(c.Request().Context(), permissionChangeRequest.UserLogin, permissionChangeRequest.NewRole); err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusInternalServerError, "failed to change role")
	}

	return c.NoContent(http.StatusOK)
}