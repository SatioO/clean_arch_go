package usecase

import (
	"context"

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
func (r *ResourceUsecase) CreateResource(context.Context, domain.Resource) error {
	return nil
}
