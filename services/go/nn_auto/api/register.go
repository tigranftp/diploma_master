package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"nn_auto/jwt_token"
	"nn_auto/passwords"
	"strings"
)

func (a *API) registerHandler() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {

		loginRAW, ok := req.Header["Login"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no login in header"))
			return
		}
		login := strings.ToLower(loginRAW[0])

		passwordRAW, ok := req.Header["Password"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no password in header"))
			return
		}
		password := passwordRAW[0]

		emailRAW, ok := req.Header["Email"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no email in header"))
			return
		}
		email := emailRAW[0]

		userDB, err := a.store.CreateUser(login, passwords.HashPassword(password), email, context.Background())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}

		jwt, refresh, err := jwt_token.GenerateTokens(userDB.UserID, userDB.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("generating tokens error"))
			return
		}
		fmt.Println(*userDB)

		err = a.store.StoreRefreshToken(userDB.UserID, refresh)
		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		bts, _ := json.Marshal(struct {
			Token   string `json:"jwt_token"`
			Refresh string `json:"refresh_token"`
		}{
			Token:   jwt,
			Refresh: refresh,
		})
		w.Write(bts)
		return
	}
}
