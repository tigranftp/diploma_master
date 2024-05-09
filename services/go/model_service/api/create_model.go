package api

import (
	"net/http"
)

func (a *API) createModelHandler() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {

		//путь по которому хранятся скрипт и файлы
		pathRAW, ok := req.Header["path"]
		if !ok || len(pathRAW) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no path in header"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}
}
