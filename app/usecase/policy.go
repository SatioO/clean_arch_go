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
	var files []domain.Policy
	var policy domain.Policy
	path := "./static/policies/"

	fileInfo, _ := ioutil.ReadDir(path)

	for _, v := range fileInfo {
		file, _ := ioutil.ReadFile(path + v.Name())
		json.Unmarshal(file, &policy)
		files = append(files, policy)
	}

	return files, nil
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
