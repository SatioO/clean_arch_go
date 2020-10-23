package usecase

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/satioO/togo/app/domain"
	"github.com/satioO/togo/app/usecase/port/repository"
)

// ResourceUsecase ...
type ResourceUsecase struct {
	repo repository.ResourceRepository
}

// NewResourceUsecase ...
func NewResourceUsecase(repo repository.ResourceRepository) *ResourceUsecase {
	return &ResourceUsecase{repo}
}

// ListResources ...
func (r *ResourceUsecase) ListResources(context.Context) []domain.Resource {
	return nil
}

// CreateResource ...
func (r *ResourceUsecase) CreateResource(ctx context.Context, resource domain.Resource) error {
	file, err := json.Marshal(&resource)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./static/resources/"+resource.Name+".json", file, 0644)
	return err
}
