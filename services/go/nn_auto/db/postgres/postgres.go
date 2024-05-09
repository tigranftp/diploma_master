package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"nn_auto/types"
)

// Postgres ...
type Postgres struct {
	db  *pgx.Conn
	dsn string
}

func New(dsn string, ctx context.Context) (*Postgres, error) {
	conn, err := pgx.Connect(ctx, "postgres://postgres:qwerty@localhost:11235/postgres")
	if err != nil {
		return nil, fmt.Errorf("could not connect to postgres: %v", err)
	}

	return &Postgres{conn, dsn}, nil
}

func (p *Postgres) Close() error {
	return p.db.Close(context.Background())
}

func (p *Postgres) CreateUser(username, passwordHash, email string, ctx context.Context) (*types.UserDB, error) {
	row := p.db.QueryRow(ctx, createUserQuery, username, passwordHash, email)

	var id int64
	err := row.Scan(&id)
	return &types.UserDB{UserID: id,
		Username:     username,
		PasswordHash: passwordHash,
		Email:        email,
	}, err
}

func (p *Postgres) StoreRefreshToken(userID int64, refreshToken string) error {
	_, err := p.db.Exec(context.Background(), setRefreshQuery, refreshToken, userID)
	return err
}

func (p *Postgres) GetUserByLogin(login string) (*types.UserDB, error) {
	row := p.db.QueryRow(context.Background(), getUserByLoginQuery, login)

	res := new(types.UserDB)

	err := row.Scan(&res.UserID, &res.Username, &res.PasswordHash, &res.Email, &res.RefreshToken)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return res, err
}

func (p *Postgres) GetUserByRefreshToken(refreshToken string) (*types.UserDB, error) {
	row := p.db.QueryRow(context.Background(), getUserByRefreshQuery, refreshToken)

	res := new(types.UserDB)

	err := row.Scan(&res.UserID, &res.Username, &res.PasswordHash, &res.Email, &res.RefreshToken)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return res, err
}
