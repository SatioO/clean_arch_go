package user

import (
	"context"

	"github.com/satioO/togo/domain"
)

// Repository ...
type Repository interface {
	FindAll(context context.Context) ([]domain.User, error)
	Save(context context.Context, user *domain.User) error
}
