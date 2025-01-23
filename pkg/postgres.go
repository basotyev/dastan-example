package pkg

import (
	"context"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresConnection(dsn string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, errors.Join(err, errors.New("could not connect to db"))
	}

	m, err := migrate.New("file://db/migrations", dsn)
	if err != nil {
		return nil, errors.Join(err, errors.New("could not create migrate struct"))
	}
	if err := m.Up(); err != nil {
		if !errors.Is(migrate.ErrNoChange, err) {
			return nil, errors.Join(err, errors.New("m.Up: could not apply migrations"))
		}
	}
	return dbPool, nil
}
