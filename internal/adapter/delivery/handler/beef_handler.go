package handler

import (
	"net/http"
	"piefiredire/internal/adapter/dto"
	"piefiredire/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type BeefHandler struct {
	service ports.BeefService
}

func NewBeefHandler(service ports.BeefService) BeefHandler {
	return BeefHandler{
		service: service,
	}
}

func (h *BeefHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/beef", h.Get)
	e.GET("/beef/summary", h.GetSummary)
}

func (h *BeefHandler) Get(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *BeefHandler) GetSummary(c echo.Context) error {
	summary, err := h.service.Summary()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	resDTO := &dto.BeefSummaryResponse{
		Beef: summary,
	}

	return c.JSON(http.StatusOK, resDTO)
}
