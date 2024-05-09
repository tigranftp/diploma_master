package api

import (
	"nn_auto/converter"
	"nn_auto/jwt_token"
	"nn_auto/types"
	"strings"
)

func (a *API) TokenMiddleware(jwtToken, refreshToken string) (*types.UserAPI, error) {
	userID, username, err := jwt_token.ParseToken(jwtToken)
	if err == nil {
		return &types.UserAPI{UserID: userID, Username: username}, err
	}
	if !strings.Contains(err.Error(), "token is expired") {
		return nil, err
	}
	userDB, err := a.store.GetUserByRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	err = jwt_token.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	return converter.ConvertUserDB(userDB), nil
}
