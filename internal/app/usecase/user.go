package usecase

import (
	"context"
	"errors"
	"newExampleServer/internal/app/service/user"
	"newExampleServer/internal/models"
	"strings"
)

type UserUseCase interface {
	Register(ctx context.Context, name, email, password string) error
}

func NewUserUseCase(userService user.Service) UserUseCase {
	return &userUseCase{
		userService: userService,
	}
}

type userUseCase struct {
	userService user.Service
}

func (u *userUseCase) Register(ctx context.Context, name, email, password string) error {
	if !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}
	if len([]byte(password)) < 5 || !strings.Contains(password, "!") {
		return errors.New("weak password")
	}
	var usr models.User
	usr.Name = name
	usr.Email = email
	usr.Password = password
	err := u.userService.CreateUser(ctx, usr)
	if err != nil {
		return err
	}
	return nil
}
