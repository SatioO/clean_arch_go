package usecase

import (
	"context"

	"github.com/satioO/togo/app/domain"
)

// ResourceUsecase ...
type ResourceUsecase interface {
	ListResources(context.Context) ([]domain.Resource, error)
	CreateResource(context.Context, domain.Resource) error
}
