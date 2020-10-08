package usecase

import (
	"context"

	"github.com/satioO/togo/app/domain"
)

// PolicyUsecase ...
type PolicyUsecase interface {
	ListPolicies(context.Context) ([]domain.Policy, error)
	CreatePolicy(context.Context, *domain.Policy) error
}
