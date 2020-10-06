package user

import (
	"context"

	"github.com/satioO/togo/domain"
)

// Usecase ...
type Usecase interface {
	FindUsers(context context.Context) ([]domain.User, error)
	SaveUser(context context.Context, user *domain.User) error
}
