package usecase

import (
	"context"

	"github.com/satioO/togo/app/domain"
	"github.com/satioO/togo/app/usecase/port/repository"
)

// UserUsecase ...
type UserUsecase struct {
	repo repository.UserRepository
}

// NewUserUsecase ...
func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo}
}

// FindUsers ...
func (u *UserUsecase) FindUsers(ctx context.Context) ([]domain.User, error) {
	return u.repo.FindAll(ctx)
}

// SaveUser ...
func (u *UserUsecase) SaveUser(ctx context.Context, user *domain.User) error {
	return u.repo.Save(ctx, user)
}
