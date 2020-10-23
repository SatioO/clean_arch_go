package usecase

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/satioO/togo/app/domain"
)

// PolicyUseCase ...
type PolicyUseCase struct{}

// NewPolicyUsecase ...
func NewPolicyUsecase() *PolicyUseCase {
	return &PolicyUseCase{}
}

// ListPolicies ....
func (r *PolicyUseCase) ListPolicies(ctx context.Context) ([]domain.Policy, error) {
	return nil, nil
}

// CreatePolicy ....
func (r *PolicyUseCase) CreatePolicy(ctx context.Context, policy *domain.Policy) error {
	file, err := json.Marshal(&policy)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./static/policies/"+policy.Name+".json", file, 0644)
	return err
}
