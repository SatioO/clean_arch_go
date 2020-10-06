package usecase

import (
	"context"

	"github.com/satioO/togo/app/domain"
)

// UserUsecase ...
type UserUsecase interface {
	FindUsers(context.Context) ([]domain.User, error)
	SaveUser(context.Context, *domain.User) error
}
