package repository

import (
	"context"

	"github.com/satioO/togo/app/domain"
)

// ResourceRepository ...
type ResourceRepository interface {
	Find(context.Context) ([]domain.Resource, error)
	Save(context.Context, *domain.Resource) error
}
