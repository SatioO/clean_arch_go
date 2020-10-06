package repository

import (
	"context"

	"github.com/satioO/togo/app/domain"
)

// UserRepository ...
type UserRepository interface {
	FindAll(context.Context) ([]domain.User, error)
	Save(context.Context, *domain.User) error
}
