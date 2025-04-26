package service

import (
	"context"
	"yuyuid.id/internal/model"
	"yuyuid.id/internal/repository"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepositoryImpl
}

func (s *UserServiceImpl) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	return s.UserRepo.FindByEmail(ctx, email)
}

func NewUserService(userRepo repository.UserRepositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{UserRepo: userRepo}
}

func (s *UserServiceImpl) Index(ctx context.Context) ([]model.User, error) {
	return s.UserRepo.FindAll(ctx)
}

func (s *UserServiceImpl) SaveUser(ctx context.Context, user *model.User) error {
	return s.UserRepo.Save(ctx, user)
}

func (s *UserServiceImpl) IsUserExist(ctx context.Context, email string) error {
	return s.UserRepo.IsExist(ctx, email)
}
