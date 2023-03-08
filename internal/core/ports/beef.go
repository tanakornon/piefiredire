package ports

import "piefiredire/internal/core/domain"

type BeefRepository interface {
	GetText() (string, error)
}

type BeefService interface {
	Summary() (domain.BeefSummary, error)
}
