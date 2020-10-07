package repository

import (
	"context"

	"github.com/satioO/togo/app/domain"
)

// RoleRepository ...
type RoleRepository interface {
	Find(context.Context) ([]domain.Role, error)
	Save(context.Context, *domain.Role) error
}
