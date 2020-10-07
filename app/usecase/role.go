package usecase

import (
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

// CreateRole ....
func (r *RoleUseCase) CreateRole() error {
	return nil
}
