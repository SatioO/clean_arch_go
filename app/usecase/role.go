package usecase

import (
	"context"

	"github.com/satioO/togo/app/domain"
	"github.com/satioO/togo/app/usecase/port/repository"
)

// RoleUseCase ...
type RoleUseCase struct {
	repo repository.RoleRepository
}

// NewRoleUsecase ...
func NewRoleUsecase(repo repository.RoleRepository) *RoleUseCase {
	return &RoleUseCase{repo}
}

// ListRoles ....
func (r *RoleUseCase) ListRoles(ctx context.Context) ([]domain.Role, error) {
	return r.repo.Find(ctx)
}

// CreateRole ....
func (r *RoleUseCase) CreateRole(ctx context.Context, role *domain.Role) error {
	return r.repo.Save(ctx, role)
}
