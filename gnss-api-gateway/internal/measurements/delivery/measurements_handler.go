package measurements_handler

import (
	measurements_domain_gateway "gnss-radar/gnss-api-gateway/internal/measurements"
	"gnss-radar/gnss-api-gateway/pkg/mwutils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type MeasurementsHandler struct {
	measurementsUsecase measurements_domain_gateway.Usecase
	logger              *logrus.Logger
}

func NewHandler(
	measurements measurements_domain_gateway.Usecase,
	logger *logrus.Logger,
) MeasurementsHandler {
	return MeasurementsHandler{
		measurementsUsecase: measurements,
		logger:              logger,
	}
}

func (h *MeasurementsHandler) GetEphemeris(c echo.Context) error {

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

	ephemeris, total, err := h.measurementsUsecase.GetEphemeris(c.Request().Context(), uint64(page), uint64(size))
	if err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusUnauthorized, "invalid data")
	}

	return c.JSON(http.StatusOK, measurements_domain_gateway.GetEphemerisResponse{
		Ephemeris: ephemeris,
		Total:     total,
		Page:      uint64(page),
	})
}

func (h *MeasurementsHandler) UploadEphemeris(c echo.Context) error {

	ctx := c.Request().Context()

	_, err := mwutils.GetUserID(ctx)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusUnauthorized, "No session id provided")
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		h.logger.Error("Load Ephemeris: failed to form file from provided name: ", err)
		return c.JSON(http.StatusBadRequest, "Failed to upload ephemeris")
	}

	file, err := fileHeader.Open()
	if err != nil {
		h.logger.Error("Load Ephemeris: failed to open file: ", err)
		return c.JSON(http.StatusBadRequest, "Failed to upload ephemeris")
	}
	defer file.Close()

	err = h.measurementsUsecase.LoadEphemeris(c.Request().Context(), file, fileHeader.Filename)
	if err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusInternalServerError, "invalid data")
	}

	return c.NoContent(http.StatusOK)
}

func (h *MeasurementsHandler) GetSatellitesPosition(c echo.Context) error {

	ctx := c.Request().Context()

	_, err := mwutils.GetUserID(ctx)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusUnauthorized, "No session id provided")
	}

	satellites, err := h.measurementsUsecase.GetSatellitesPosition(ctx)
	if err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, satellites)
}

func (h *MeasurementsHandler) GetSatellitesIntervals(c echo.Context) error {

	ctx := c.Request().Context()

	_, err := mwutils.GetUserID(ctx)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusUnauthorized, "No session id provided")
	}

	getIntervals := measurements_domain_gateway.GetSatellitesIntervalsRequest{}
	if err := c.Bind(&getIntervals); err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusBadRequest, "failed to parse request data")
	}

	satellites, err := h.measurementsUsecase.GetSatellitesIntervals(ctx, getIntervals.StartDatime, getIntervals.EndDatetime, getIntervals.Satellites)
	if err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, satellites)
}
