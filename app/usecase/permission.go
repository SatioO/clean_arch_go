package usecase

import (
	"context"

	"github.com/satioO/togo/app/domain"
	"github.com/satioO/togo/app/usecase/port/repository"
)

// PermissionUseCase ...
type PermissionUseCase struct {
	repo repository.PermissionRepository
}

// NewPermissionUsecase ...
func NewPermissionUsecase(repo repository.PermissionRepository) *PermissionUseCase {
	return &PermissionUseCase{repo}
}

// ListPermissions ....
func (r *PermissionUseCase) ListPermissions(ctx context.Context) ([]domain.Permission, error) {
	return r.repo.Find(ctx)
}

// CreatePermission ....
func (r *PermissionUseCase) CreatePermission(ctx context.Context, permission *domain.Permission) error {
	return r.repo.Save(ctx, permission)
}
