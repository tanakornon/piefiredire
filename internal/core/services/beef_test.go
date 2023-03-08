package services_test

import (
	"errors"
	"piefiredire/internal/core/services"
	"piefiredire/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeefSummary(t *testing.T) {
	repo := new(mocks.BeefRepositoryMock)
	service := services.NewBeefService(repo)

	text := "t-bone-bone t-bone"

	repo.On("GetText").Return(text, nil)

	result, err := service.Summary()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, 1, result["t-bone-bone"])
	assert.Equal(t, 1, result["t-bone"])
}

func TestBeefSummary_NotFound(t *testing.T) {
	repo := new(mocks.BeefRepositoryMock)
	service := services.NewBeefService(repo)

	text := ". . . ! ! !"

	repo.On("GetText").Return(text, nil)

	result, err := service.Summary()

	assert.NoError(t, err)
	assert.Equal(t, 0, len(result))
}

func TestBeefSummary_DuplicateWord(t *testing.T) {
	repo := new(mocks.BeefRepositoryMock)
	service := services.NewBeefService(repo)

	text := "t-bone-bone-.. . ..-t-bone-bone"

	repo.On("GetText").Return(text, nil)

	result, err := service.Summary()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, 2, result["t-bone-bone"])
}

func TestBeefSummary_Error(t *testing.T) {
	repo := new(mocks.BeefRepositoryMock)
	service := services.NewBeefService(repo)

	text := ""
	errMsg := errors.New("error message")

	repo.On("GetText").Return(text, errMsg)

	_, err := service.Summary()

	assert.Error(t, err)
}
