package usecase

import (
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

// ListPermission ....
func (r *PermissionUseCase) ListPermission() error {
	return nil
}

// CreatePermission ....
func (r *PermissionUseCase) CreatePermission() error {
	return nil
}
