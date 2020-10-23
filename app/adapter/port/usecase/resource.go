package usecase

import (
	"context"

	"github.com/satioO/togo/app/domain"
)

// ResourceUsecase ...
type ResourceUsecase interface {
	ListResources(context.Context) []domain.Resource
	CreateResource(context.Context, domain.Resource) error
}
