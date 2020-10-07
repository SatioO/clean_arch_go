package usecase

import (
	"context"

	"github.com/satioO/togo/app/domain"
)

// PermissionUsecase ...
type PermissionUsecase interface {
	ListPermissions(context.Context) ([]domain.Permission, error)
	CreatePermission(context.Context, *domain.Permission) error
}
