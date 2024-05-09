package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nn_auto/jwt_token"
	"nn_auto/passwords"
	"strings"
)

func (a *API) loggingHandler() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		passwordRAW, ok := req.Header["Password"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no password in header"))
			return
		}
		loginRAW, ok := req.Header["Login"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no login in header"))
			return
		}
		login := strings.ToLower(loginRAW[0])
		password := passwordRAW[0]

		userDB, err := a.store.GetUserByLogin(login)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			w.Write([]byte("something wrong with storage"))
			return
		}
		if userDB == nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(userDB)
			w.Write([]byte("login or password is incorrect"))
			return
		}
		ok = passwords.CheckPassword([]byte(userDB.PasswordHash), password)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("login or password is incorrect"))
			return
		}
		jwt, refresh, err := jwt_token.GenerateTokens(userDB.UserID, userDB.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("generating tokens error"))
			return
		}

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
