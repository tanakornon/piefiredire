package services

import (
	"piefiredire/internal/core/domain"
	"piefiredire/internal/core/ports"
	"regexp"
)

type beefService struct {
	repo ports.BeefRepository
}

func NewBeefService(repo ports.BeefRepository) ports.BeefService {
	return beefService{
		repo: repo,
	}
}

// Summary implements ports.BeefService
func (s beefService) Summary() (domain.BeefSummary, error) {
	text, err := s.repo.GetText()

	if err != nil {
		return nil, err
	}

	regex := regexp.MustCompile(`\w+(?:-\w+)*`)
	tokens := regex.FindAllString(text, -1)

	summary := domain.BeefSummary{}
	for _, token := range tokens {
		summary[token]++
	}

	return summary, nil
}
