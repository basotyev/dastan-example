package user

import (
	"context"
	"newExampleServer/internal/app/service/user/repo"
	"newExampleServer/internal/models"
)

type Service interface {
	CreateUser(context.Context, models.User) error
	GetUserById(context.Context, int) (*models.User, error)
	GetUserByEmail(context.Context, string) (*models.User, error)
	UpdateUser(context.Context, *models.User) error
	DeleteUserById(context.Context, int) error
}

type service struct {
	repo repo.Repository
}

func (s *service) CreateUser(ctx context.Context, user models.User) error {
	return s.repo.CreateUser(ctx, user)
}

func (s *service) GetUserById(ctx context.Context, i int) (*models.User, error) {
	return s.repo.GetUserById(ctx, i)
}

func (s *service) GetUserByEmail(ctx context.Context, s2 string) (*models.User, error) {
	return s.repo.GetUserByEmail(ctx, s2)

}

func (s *service) UpdateUser(ctx context.Context, user *models.User) error {
	return s.repo.UpdateUser(ctx, user)
}

func (s *service) DeleteUserById(ctx context.Context, i int) error {
	return s.repo.DeleteUserById(ctx, i)
}

func New(repo repo.Repository) Service {
	return &service{repo: repo}
}
