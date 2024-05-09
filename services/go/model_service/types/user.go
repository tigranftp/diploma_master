package types

type UserAPI struct {
	UserID   int64
	Username string
	Email    string
}

type UserDB struct {
	UserID       int64
	Username     string
	PasswordHash string
	Email        string
	RefreshToken *string
}
