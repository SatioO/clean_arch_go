package usecase

import (
	"context"

	"github.com/satioO/togo/app/domain"
)

// RoleUsecase ...
type RoleUsecase interface {
	ListRoles(context.Context) ([]domain.Role, error)
	CreateRole(context.Context, *domain.Role) error
}
