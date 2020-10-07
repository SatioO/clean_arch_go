package repository

import (
	"context"

	"github.com/satioO/togo/app/domain"
)

// PermissionRepository ...
type PermissionRepository interface {
	Find(context.Context) ([]domain.Permission, error)
	Save(context.Context, *domain.Permission) error
}
