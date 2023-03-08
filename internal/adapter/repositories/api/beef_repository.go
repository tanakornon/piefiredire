package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"piefiredire/internal/core/ports"
)

type beefRepository struct {
	baseUrl string
}

func NewBeefRepository(baseUrl string) ports.BeefRepository {
	return beefRepository{
		baseUrl: baseUrl,
	}
}

func (repo beefRepository) GetText() (string, error) {
	requestURL := repo.baseUrl
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New("not ok")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}
