package dto

import "piefiredire/internal/core/domain"

type BeefSummaryRequest struct {
	Type   string `query:"type"`
	Paras  string `query:"paras"`
	Format string `query:"format"`
}

type BeefSummaryResponse struct {
	Beef domain.BeefSummary `json:"beef"`
}
