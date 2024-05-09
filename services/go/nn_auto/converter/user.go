package converter

import (
	"nn_auto/types"
)

func ConvertUserDB(user *types.UserDB) *types.UserAPI {
	return &types.UserAPI{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
	}
}
