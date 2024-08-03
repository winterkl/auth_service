package auth_repo

import (
	"auth/internal/app_errors"
	auth_entity "auth/internal/domain/auth/entity"
	"auth/pkg/postgres"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
)

type Repo struct {
	db *postgres.Postgres
}

func New(db *postgres.Postgres) *Repo {
	return &Repo{db}
}

func (r *Repo) CreateUser(ctx context.Context, user auth_entity.User) (int, error) {
	if _, err := r.db.NewInsert().
		Model(&user).
		Exec(ctx); err != nil {
		if pgErr := err.(*pgconn.PgError); pgErr.Code == r.db.Errors.CodeUniqueConstraint {
			return 0, &app_errors.UserAlreadyExists{Login: user.Login}
		}
		return 0, fmt.Errorf("NewInsert: %w", err)
	}
	return user.ID, nil
}

func (r *Repo) GetUserByAuthData(ctx context.Context, user auth_entity.User) error {
	if err := r.db.NewSelect().
		Model(&user).
		Where("login = ?", user.Login).
		Where("password = ?", user.Password).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &app_errors.UserNotFound{}
		}
		return fmt.Errorf("NewSelect: %w", err)
	}
	return nil
}
