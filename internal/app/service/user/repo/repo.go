package repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"newExampleServer/internal/models"
)

type Repository interface {
	CreateUser(context.Context, models.User) error
	GetUserById(context.Context, int) (*models.User, error)
	GetUserByEmail(context.Context, string) (*models.User, error)
	UpdateUser(context.Context, *models.User) error
	DeleteUserById(context.Context, int) error
}

type repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user models.User) error {
	_, err := r.db.Exec(ctx, `INSERT INTO users(name, email, password) VALUES ($1, $2, $3)`, user.Name, user.Email, user.Password)
	return err
}

func (r *repository) GetUserById(ctx context.Context, id int) (*models.User, error) {
	var usr models.User
	err := r.db.QueryRow(ctx, `SELECT id, name, email, password, created_at, updated_at from users WHERE id = $1`, id).Scan(&usr.Id, &usr.Name, &usr.Email, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var usr models.User
	err := r.db.QueryRow(ctx, `SELECT id, name, email, password, created_at, updated_at from users WHERE email ilike $1`, email).Scan(&usr.Id, &usr.Name, &usr.Email, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func (r *repository) UpdateUser(ctx context.Context, user *models.User) error {
	_, err := r.db.Exec(ctx, `UPDATE users set name = $1, email = $2, password = $3 WHERE id = $4`, user.Name, user.Email, user.Password, user.Id)
	return err
}

func (r *repository) DeleteUserById(ctx context.Context, i int) error {
	res, err := r.db.Exec(ctx, `DELETE FROM users where id = $1`, i)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("not found")
	}
	return nil
}
