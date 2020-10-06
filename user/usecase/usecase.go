package usecase

import (
	"context"

	"github.com/satioO/togo/app/user"
	"github.com/satioO/togo/domain"
)

// UserUsecase ...
type UserUsecase struct {
	repo user.Repository
}

// NewUserUsecase ...
func NewUserUsecase(repo user.Repository) *UserUsecase {
	return &UserUsecase{
		repo,
	}
}

// FindUsers ....
func (u *UserUsecase) FindUsers(ctx context.Context) ([]domain.User, error) {
	return u.repo.FindAll(ctx)
}

// SaveUser ....
func (u *UserUsecase) SaveUser(ctx context.Context, user *domain.User) error {
	return u.repo.Save(ctx, user)
}
