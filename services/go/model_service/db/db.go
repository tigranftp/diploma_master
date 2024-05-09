package db

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"nn_auto/types"
)

type Storage interface {
	Close() error
	StoreRefreshToken(userID int64, refreshToken string) error
	CreateUser(username, password, email string, ctx context.Context) (*types.UserDB, error)
	GetUserByLogin(login string) (*types.UserDB, error)
	GetUserByRefreshToken(refreshToken string) (*types.UserDB, error)
}
