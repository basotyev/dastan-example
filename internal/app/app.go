package app

import (
	"github.com/jackc/pgx/v4/pgxpool"
	userS "newExampleServer/internal/app/service/user"
	"newExampleServer/internal/app/service/user/repo"
	"newExampleServer/internal/app/usecase"
)

type DI struct {
	UseCases UseCases
}

type UseCases struct {
	UserUC usecase.UserUseCase
}

func NewDI(db *pgxpool.Pool) *DI {
	userRepo := repo.New(db)
	userService := userS.New(userRepo)
	userUC := usecase.NewUserUseCase(userService)

	return &DI{
		UseCases: UseCases{
			UserUC: userUC,
		},
	}
}
