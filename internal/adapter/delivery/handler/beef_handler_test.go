package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"piefiredire/internal/adapter/delivery/handler"
	"piefiredire/internal/adapter/dto"
	"piefiredire/internal/core/services"
	"piefiredire/mocks"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetSummary_OK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	r := new(mocks.BeefRepositoryMock)
	s := services.NewBeefService(r)
	h := handler.NewBeefHandler(s)
	e := echo.New()
	c := e.NewContext(req, rec)

	text := "t-bone-bone t-bone"

	r.On("GetText").Return(text, nil)

	dto := dto.BeefSummaryResponse{
		Beef: map[string]int{
			"t-bone-bone": 1,
			"t-bone":      1,
		},
	}

	bytes, _ := json.Marshal(dto)
	expected := fmt.Sprintf("%s\n", bytes)

	if assert.NoError(t, h.GetSummary(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestGetSummary_Error(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	r := new(mocks.BeefRepositoryMock)
	s := services.NewBeefService(r)
	h := handler.NewBeefHandler(s)
	e := echo.New()
	c := e.NewContext(req, rec)

	text := ""
	errMsg := "error message"

	r.On("GetText").Return(text, errors.New(errMsg))

	dto := dto.ErrorResponse{
		Message: errMsg,
	}

	bytes, _ := json.Marshal(dto)
	expected := fmt.Sprintf("%s\n", bytes)

	if assert.NoError(t, h.GetSummary(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
