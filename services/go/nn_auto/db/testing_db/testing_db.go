package testing_db

import (
	"nn_auto/types"
)

type TestingDB struct {
}

// NewTestingDB ...
func NewTestingDB() *TestingDB {

	return &TestingDB{}
}

// Open ...
func (d *TestingDB) Open() error {

	return nil
}

// StoreRefreshToken ...
func (d *TestingDB) StoreRefreshToken(userID int64, refreshToken string) error {

	return nil
}

// GetUserByLogin ...
func (d *TestingDB) GetUserByLogin(login string) (*types.UserDB, error) {

	return &types.UserDB{
		UserID:       1,
		Username:     "tigranFTP",
		PasswordHash: string([]byte{83, 89, 138, 227, 233, 137, 3, 230, 127, 114, 246, 232, 115, 172, 145, 126, 25, 172, 56, 191, 71, 110, 65, 5, 220, 70, 90, 86, 113, 108, 194, 209, 191, 66, 177, 32, 238, 142, 66, 234}),
		Email:        "tigranftp@gmail.com",
		RefreshToken: "",
	}, nil
}

// GetUserByRefreshToken ...
func (d *TestingDB) GetUserByRefreshToken(refreshToken string) (*types.UserDB, error) {

	return nil, nil
}

// Close ...
func (d *TestingDB) Close() error {
	return nil
}
