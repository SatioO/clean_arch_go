package usecase

import (
	"context"

	"github.com/satioO/togo/app/domain"
	"github.com/satioO/togo/app/usecase/port/repository"
)

// PolicyUseCase ...
type PolicyUseCase struct {
	repo repository.PolicyRepository
}

// NewPolicyUsecase ...
func NewPolicyUsecase(repo repository.PolicyRepository) *PolicyUseCase {
	return &PolicyUseCase{repo}
}

// ListPolicies ....
func (r *PolicyUseCase) ListPolicies(ctx context.Context) ([]domain.Policy, error) {
	return r.repo.Find(ctx)
}

// CreatePolicy ....
func (r *PolicyUseCase) CreatePolicy(ctx context.Context, policy *domain.Policy) error {
	return r.repo.Save(ctx, policy)
}
