package usecase

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/satioO/togo/app/domain"
)

// ResourceUsecase ...
type ResourceUsecase struct{}

// NewResourceUsecase ...
func NewResourceUsecase() *ResourceUsecase {
	return &ResourceUsecase{}
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
