package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

const (
	_errorCodeConstraint       = "23503"
	_errorCodeUniqueConstraint = "23505"
)

type ErrorsCode struct {
	CodeConstraint       string
	CodeUniqueConstraint string
}

type Postgres struct {
	*bun.DB
	Errors ErrorsCode
}

func New(user, password, host, dbName, sslMode string, port int) (*Postgres, error) {

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%s", user, password, host, port, dbName, sslMode)
	config, err := pgx.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	config.PreferSimpleProtocol = true

	sqlDB := stdlib.OpenDB(*config)

	return &Postgres{
		Errors: ErrorsCode{
			CodeConstraint:       _errorCodeConstraint,
			CodeUniqueConstraint: _errorCodeUniqueConstraint,
		},
		DB: bun.NewDB(sqlDB, pgdialect.New()),
	}, nil
}
