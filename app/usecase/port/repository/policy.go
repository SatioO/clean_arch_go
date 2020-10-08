package repository

import (
	"context"

	"github.com/satioO/togo/app/domain"
)

// PolicyRepository ...
type PolicyRepository interface {
	Find(context.Context) ([]domain.Policy, error)
	Save(context.Context, *domain.Policy) error
}
